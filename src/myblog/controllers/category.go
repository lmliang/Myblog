package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("catename")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	var err error

	this.TplNames = "category.html"
	this.Data["IsCategory"] = true
	this.Data["Categorys"], err = models.GetAllCategorys()
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	if err != nil {
		beego.Error(err)
	}
}
