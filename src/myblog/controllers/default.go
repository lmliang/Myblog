package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	curCate := this.Input().Get("category")
	curLabel := this.Input().Get("label")

	this.TplNames = "home.html"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["Topics"], _ = models.GetAllTopics(curCate, curLabel, true)
	this.Data["Categorys"], _ = models.GetAllCategorys()
}
