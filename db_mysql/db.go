package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "大一下学期/github.com/go-sql-driver/mysql"
)

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
//func AddUser(u models.User)(int64,error){
//	md5Hash := md5.New()
//	md5Hash.Write([]byte(u.Password))
//	passwordBytes := md5Hash.Sum(nil)
//	u.Password = hex.EncodeToString(passwordBytes)
//	result,err := Db.Exec("insert into user(phone,password) values(?,?)",u.Phone,u.Password)
//	if err != nil{
//		fmt.Println(err.Error())
//		return -1,err
//	}
//	row, err := result.RowsAffected()
//	if err != nil{
//		fmt.Println(err.Error())
//		return -1, err
//	}
//	return row, nil
//}
