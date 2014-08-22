package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")

		this.Redirect("/", 301)
		return
	}

	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autologin := this.Input().Get("autologin") == "on"

	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autologin {
			maxAge = 3600
		}

		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")

		this.Redirect("/", 301)
		return
	} else {
		this.Redirect("/login", 301)
	}
}

func checkAccount(ctx *context.Context) bool {
	name, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}

	username := name.Value

	pwd, err := ctx.Request.Cookie("password")
	if err != nil {
		return false
	}

	password := pwd.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
