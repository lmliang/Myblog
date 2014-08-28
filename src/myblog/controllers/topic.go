package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.Data["Topics"], _ = models.GetAllTopics("", "", false)
	this.TplNames = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}

	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	labels := this.Input().Get("labels")

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		attachment = fh.Filename
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content, category, labels, attachment)
	} else {
		err = models.ModifyTopic(tid, title, content, category, labels, attachment)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.TplNames = "topic_add.html"
}

func (this *TopicController) View() {
	id := this.Ctx.Input.Params["0"]
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	replys, err := models.GetTopicComments(id)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Tid"] = id
	this.Data["Topic"] = topic
	this.Data["Replys"] = replys
	this.Data["Labels"] = strings.Split(topic.Labels, " ")
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplNames = "topic_view.html"
}

func (this *TopicController) Modify() {
	id := this.Input().Get("tid")
	topic, err := models.GetTopic(id)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Tid"] = id
	this.Data["Topic"] = topic
	this.TplNames = "topic_modify.html"
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}

	tid := this.Ctx.Input.Params["0"]

	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}
