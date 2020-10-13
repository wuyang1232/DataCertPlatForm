package controllers

import "beego"

type HomeController struct {
	beego.Controller
}
func (r *HomeController)Post(){
	r.TplName = "login.html"
}
