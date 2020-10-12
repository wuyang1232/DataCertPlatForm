package controllers

import "beego"

type IndexNew struct {
	beego.Controller
}
func (r *IndexNew)Post(){
	r.TplName = "login.html"
}
