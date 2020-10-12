package controllers

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
//该方法用于处理用户注册的逻辑
func (r *RegisterController)Get(){

	r.TplName = "login.html"
	//如果失败提示错误信息
}
func (r *RegisterController) Post(){
	//1、解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("抱歉...数据解析失败，请重试!")
		return
	}
	//2、将解析到的数据保存到数据库中
	if len(user.Phone) == 11{
		if len(user.Password) >= 6 && len(user.Password) <= 18{
			row, err := db_mysql.AddUser(user)
			if err != nil{
				r.Ctx.WriteString("数据导入数据库出错，请重试!")
			}
			fmt.Println(row)
			//3、将处理结果返回给客户端浏览器
			//如果成功，跳转登录页面
			r.TplName = "login.html"
		}else if len(user.Password) > 18|| len(user.Password) < 6{
			r.Ctx.WriteString("密码长度错误，请重新输入")
		}
	}else {
		//如果失败提示错误信息
		r.Ctx.WriteString("电话号码错误")
	}
}
