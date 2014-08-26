package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	nick := this.Input().Get("nickname")
	content := this.Input().Get("content")

	err := models.AddReply(tid, nick, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	rid := this.Input().Get("id")

	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}
