package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
	"time"
)

type LoginSmsController struct {
	beego.Controller
}
//浏览器发起的短信验证码登录请求
func (l *LoginSmsController)Get(){
	l.TplName = "login_sms.html"
}
//短信验证码登录功能
func (l *LoginSmsController)Post(){
	var smsLogin models.SmsLogin
	err := l.ParseForm(&smsLogin)
	if err != nil{
		l.Ctx.WriteString("抱歉，验证码登录数据解析失败")
		return
	}
	//1、先拿手机号查询user表，看用户是否已存在
	user := models.User{
		Phone: smsLogin.Phone,
	}
	_, err = user.QueryUserByPhone()
	if err != nil{
		l.Ctx.WriteString("该手机号还未注册，请先注册")
		return
	}

	//拿用户提交的登录信息到数据库中查询
	sms, err := models.QuerySmsRecord(smsLogin.BizId,smsLogin.Phone,smsLogin.Code)
	if err != nil{
		l.Ctx.WriteString("抱歉，验证码登录遇到错误，")
		return
	}
	if sms.BizId == ""{//验证码错误，或者手机号错误
		l.Ctx.WriteString("手机号或者验证码错误，请重新输入")
		return
	}
	now := time.Now().Unix()
	if (now - sms.TimeStamp) > 300000{
		l.Ctx.WriteString("验证码已失效，请重新获取")
		return
	}
	//登录正常，跳转主页面
	l.Data["Phone"] = smsLogin.Phone
	l.TplName = "hone.html"
}