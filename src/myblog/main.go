package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/controllers"
	"myblog/models"
	_ "myblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Run()
}
