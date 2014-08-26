package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.Data["Topics"], _ = models.GetAllTopics(false)
	this.TplNames = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.TplNames = "topic_add.html"
}

func (this *TopicController) View() {
	id := this.Input().Get("id")
	this.Data["Topic"], _ = models.GetTopic(id)
	this.TplNames = "topic_view.html"
}
