package db_mysql

import (
	"DataCertPlatform/models"
	_ "大一下学期/github.com/go-sql-driver/mysql"
)

//func QueryUser(user models.User)(int,error){
//	row := Db.QueryRow("select count(phone) admin_num from user where phone = ? and password = ?",user.Phone,user.Password)
//	var admin_num int
//	err := row.Scan(&admin_num)
//	if err != nil{
//		return 0,err
//	}
//	return admin_num,nil
//}
func QueryUser(name string)([]models.User,error){

}
