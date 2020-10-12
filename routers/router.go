package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.RegisterController{})//用户注册接口
    beego.Router("/login",&controllers.QueryUser{})
    beego.Router("/index",&controllers.IndexNew{})
}
