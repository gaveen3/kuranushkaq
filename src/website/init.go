package website

import (
	"fmt"
	"log"
	rcMw "middleware"
	"time"
	"utils"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/adaptors/websocket"
)

var (
	app = iris.New()
)

func init() {
	app.Config.Gzip = true
	app.Config.Charset = "UTF-8"
	app.Adapt(iris.DevLogger())

	app.Adapt(httprouter.New())

	app.UseGlobal(rcMw.NewServeHeader("rcs/v1.2", "realclouds.org"))
	app.UseGlobal(rcMw.NewOSEnv())

	app.Adapt(view.HTML(utils.GetProjectDir()+"/views/default", ".html").Delims("{%", "%}").Reload(true))

	ws := websocket.New(websocket.Config{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		Endpoint:         "/ws/image",
		ClientSourcePath: "/ws/realclouds_ws.js",
	})

	ws.OnConnection(imageHandle)

	app.Adapt(ws)

	app.StaticWeb("/pub", utils.GetProjectDir()+"/assets")

	vmImagesDir := utils.GetENV("VM_IMAGES_DIR")
	if len(vmImagesDir) == 0 {
		vmImagesDir = "/opt/images/vm"
	}

	app.StaticServe(vmImagesDir, "/images/vm")

	app.Adapt(sessions.New(sessions.Config{
		Cookie:                      "JSESSIONID",
		Expires:                     time.Minute * 30,
		CookieLength:                32,
		DisableSubdomainPersistence: false,
	}))

	app.Adapt(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("404.html", nil)
	})

	app.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("500.html", nil)
	})
}

// None registers an "offline" route
// see context.ExecRoute(routeName),
// iris.Default.None(...) and iris.Default.SetRouteOnline/SetRouteOffline
// For more details look: https://github.com/kataras/iris/issues/585
//
// Example: https://github.com/iris-contrib/examples/tree/master/route_state
func None(path string, handlersFn ...iris.HandlerFunc) {
	app.None(path, handlersFn...)
}

// Get registers a route for the Get http method
func Get(path string, handlersFn ...iris.HandlerFunc) {
	app.Get(path, handlersFn...)
}

// Post registers a route for the Post http method
func Post(path string, handlersFn ...iris.HandlerFunc) {
	app.Post(path, handlersFn...)
}

// Put registers a route for the Put http method
func Put(path string, handlersFn ...iris.HandlerFunc) {
	app.Put(path, handlersFn...)
}

// Delete registers a route for the Delete http method
func Delete(path string, handlersFn ...iris.HandlerFunc) {
	app.Delete(path, handlersFn...)
}

// Connect registers a route for the Connect http method
func Connect(path string, handlersFn ...iris.HandlerFunc) {
	app.Connect(path, handlersFn...)
}

// Head registers a route for the Head http method
func Head(path string, handlersFn ...iris.HandlerFunc) {
	app.Head(path, handlersFn...)
}

// Options registers a route for the Options http method
func Options(path string, handlersFn ...iris.HandlerFunc) {
	app.Options(path, handlersFn...)
}

// Patch registers a route for the Patch http method
func Patch(path string, handlersFn ...iris.HandlerFunc) {
	app.Patch(path, handlersFn...)
}

// Trace registers a route for the Trace http method
func Trace(path string, handlersFn ...iris.HandlerFunc) {
	app.Trace(path, handlersFn...)
}

// Any registers a route for ALL of the http methods (Get,Post,Put,Head,Patch,Options,Connect,Delete)
func Any(path string, handlersFn ...iris.HandlerFunc) {
	app.Any(path, handlersFn...)
}

//Run WebSite
func Run() {
	port := utils.GetENV("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	tlsPort := utils.GetENV("TLS_PORT")
	if len(tlsPort) == 0 {
		tlsPort = "8443"
	}

	tls := utils.GetENV("TLS")
	if len(tls) > 0 && "true" == tls {
		cert := utils.GetENV("CERT")
		key := utils.GetENV("KEY")
		if len(cert) > 0 && len(key) > 0 {
			app.ListenTLS(":"+tlsPort, cert, key)
		} else {
			tlsHost := utils.GetENV("TLS_HOST")
			if len(tlsHost) != 0 {
				app.ListenLETSENCRYPT(tlsHost + ":" + tlsPort)
			} else {
				log.Fatalln(fmt.Errorf("%s", "Invalid certificate configuration."))
			}
		}
	} else {
		app.Listen(":" + port)
	}
}
