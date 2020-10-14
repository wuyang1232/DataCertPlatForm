package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})//用户注册接口
    //用户登录接口
    beego.Router("/login",&controllers.LoginController{})
    //用户上传文件的功能
    beego.Router("/index",&controllers.UploadController{})
}
