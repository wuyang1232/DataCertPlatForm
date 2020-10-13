package controllers

import (
	//"DataCertPlatform/db_mysql"
	"DataCertPlatform/models"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
var Db *sql.DB

func Connect(){
	//项目配置
	config := beego.AppConfig //定义config变量，接受并赋值为全局配置变量
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称：", appName)
	port, err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目信息解析错误，请检验后重试")
	}
	fmt.Println("应用监听端口：", port)

	driver := config.String("db_driver")
	dbUser := config.String("db_root")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//1、连接数据库
	db, err := sql.Open(driver, dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	//sql.Open("mysql","root:281511@tcp(127.0.0.1:3306)/hero_lol?charset=utf8")
	if err != nil { //err 不等于nil表示连接数据库的时候出现错误，程序就在此中断，不用在往下执行
		//早发现，早解决
		panic("数据库连接失败") //panic：是指程序进入一种恐慌状态，程序会终止执行
	}
	Db = db
	fmt.Println(db)
	fmt.Println("数据库连接成功")
}
func (r *LoginController) Post(){
	//1、解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)//地址读取
	fmt.Println(user)
	fmt.Println("逗比")
	if err != nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉...用户登录信息数据解析失败，请重试!")
		return
	}
	fmt.Println("逗比1")

	//2、根据解析到的数据，执行数据库查新操作
	u,err := user.QueryUser()

	//3、判断数据库查询结果
	if err != nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉...用户登录失败，请重试!")
		return
	}

	//4、根据查询结果，返回客户端相应的信息或页面跳转
	r.Data["Phone"] = u.Phone//动态数据设置
	r.TplName = "index.html" //用户存在转入主页面
}

