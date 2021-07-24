package website

import (
	"ldap"
	"strings"

	"utils"

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

	LogDebugln("User:", username, password)

	result := Result{}

	ldapAddr := getENV("LDAP_ADDR")
	if len(ldapAddr) == 0 {
		ldapAddr = "10.0.99.100"
	}

	ldapPort := getENV("LDAP_PORT")
	if len(ldapPort) == 0 {
		ldapPort = "636"
	}

	lPort, err := utils.StringUtils(ldapPort).Int()
	if nil != err {
		LogErrorln("Invalid LDAP port. Default:636")
		lPort = 636
	}

	ldapDC := getENV("LDAP_DC")
	if len(ldapDC) == 0 {
		ldapDC = "dc=ronglian,dc=com"
	}

	ldapType := getENV("LDAP_TYPE")
	if len(ldapType) == 0 {
		ldapType = "ad"
	}

	lc := ldap.NewLDAPClient(ldapAddr, int(lPort), true, ldapDC, username, password, ldapType)
	defer lc.Close()

	if _, err := lc.Authenticate(); nil != err {
		LogDebugln(err)
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
