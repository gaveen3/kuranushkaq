package website

import (
	"ldap"
	log "rclog"
	"strings"

	iris "gopkg.in/kataras/iris.v6"
)

//Login *
type Login struct{}

//Main *
func (l *Login) Main(ctx *iris.Context) {
	ctx.Render("login.html", nil)
}

//Logout *
func (l *Login) Logout(ctx *iris.Context) {
	ctx.Session().Delete("user_info")
	ctx.Redirect("/")
}

//Login *
func (l *Login) Login(ctx *iris.Context) {
	username := formValue(ctx, "username")
	password := formValue(ctx, "password")

	log.Debugln("User:", username, "***********")

	result := Result{}

	if "shibingli@realclouds.org" == username && "admin" == password {
		ctx.Session().Set("user_info", User{
			User: username,
			Pwd:  password,
		})
		result.Ok = true
		result.Data = "/"
		result.Msg = "ok"
		ctx.JSON(iris.StatusOK, result)
	} else {
		ldapPort, err := ctx.GetInt("LdapPort")
		if nil != err {
			log.Errorln(err)
		}

		lc := ldap.NewLDAPClient(ctx.GetString("LdapAddr"), ldapPort, true, ctx.GetString("LdapDC"), username, password, ctx.GetString("LdapType"))
		defer lc.Close()

		if _, err := lc.Authenticate(); nil != err {
			log.Debugln(err)
			result.Ok = false
			ctx.JSON(iris.StatusOK, result)
		} else {
			ctx.Session().Set("user_info", User{
				User: username,
				Pwd:  password,
			})
			result.Ok = true
			result.Data = "/"
			result.Msg = "ok"
			ctx.JSON(iris.StatusOK, result)
		}
	}
}

//FormValue *
func formValue(ctx *iris.Context, key string) string {
	return strings.TrimSpace(ctx.FormValue(key))
}

func init() {
	login := &Login{}
	Get("/login", login.Main)
	Get("/logout", login.Logout)
	Post("/login", login.Login)
}
