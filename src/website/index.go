package website

import (
	"models"
	"os"
	"path/filepath"
	log "rclog"
	"strings"

	"github.com/go-fsnotify/fsnotify"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/websocket"
)

//Images *
var Images []*models.Image

//Index *
type Index struct{}

//Main *
func (i *Index) Main(ctx *iris.Context) {
	sessionUser := ctx.Session().Get("user_info")
	if nil != sessionUser {
		ctx.Render("index.html", sessionUser)
	} else {
		ctx.Redirect("/login")
	}
}

const (
	//ImageRoom *
	ImageRoom string = "iRoom"
)

//imageHandle WsImages WebSocket handle
func imageHandle(c websocket.Connection) {
	log.Debugln("Connection:", c.ID())

	imagesDir := c.Context().GetString("VMImagesDir")

	log.Debugln("VM images dir:", imagesDir)

	c.Join(ImageRoom)

	c.OnDisconnect(func() {
		newWSResult(c, "error", false, "no", "\nError: Client disconnect.")
	})

	go watcherImagesDir(c, imagesDir)

	c.On("loadImages", func(s string) {
		if err := LoadImageDir(imagesDir); nil != err {
			newWSResult(c, "error", false, "no", err.Error())
		}
		newWSResult(c, "imagesResult", true, "ok", Images)
	})
}

func watcherImagesDir(c websocket.Connection, imagesDir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Errorln(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Debugln("Watcher event:", event)
				eop := event.Op
				if eop&fsnotify.Create == fsnotify.Create || eop&fsnotify.Remove == fsnotify.Remove || eop&fsnotify.Rename == fsnotify.Rename {
					if err := LoadImageDir(imagesDir); nil != err {
						newWSResult(c, "error", false, "no", err.Error())
					}
					newWSResult(c, "imagesResult", true, "ok", Images)
				}
			case err := <-watcher.Errors:
				log.Debugln("error:", err)
			}
		}
	}()

	err = watcher.Add(imagesDir)
	if err != nil {
		log.Errorln(err)
	}
	<-done
}

func newWSResult(c websocket.Connection, event string, ok bool, msg string, data interface{}) {
	if err := c.To(ImageRoom).Emit(event, Result{
		Ok:   ok,
		Msg:  strings.TrimSpace(msg),
		Data: data,
	}); nil != err {
		log.Errorln(err)
	}
}

//LoadImageDir *
func LoadImageDir(dir string) error {
	Images = make([]*models.Image, 0, 0)
	if err := filepath.Walk(strings.TrimRight(dir, "/"), func(p string, f os.FileInfo, err error) error {
		return imagePathWalk(p, f, err)
	}); nil != err {
		return err
	}
	return nil
}

func imagePathWalk(p string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	} else if f.IsDir() {
		return nil
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		return nil
	}

	if f.Size() > 0 {
		ext := strings.ToLower(filepath.Ext(p))

		if ".iso" == ext || ".qcow2" == ext || ".raw" == ext || ".ova" == ext {

			var t string

			switch ext {
			case ".iso":
				t = "iso"
			case ".qcow2":
				t = "qcow2"
			case ".raw":
				t = "raw"
			case ".ova":
				t = "ova"
			}

			name := f.Name()
			image := &models.Image{
				Name: strings.TrimRight(name, "."+t),
				Type: t,
				Size: f.Size(),
				Path: "/images/vm/" + name,
			}
			Images = append(Images, image)
		}
	}

	return err
}

func init() {
	index := &Index{}
	Get("/", index.Main)
}
