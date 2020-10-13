package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
	_"大一下学期/github.com/go-sql-driver/mysql"
)

//import (
//	"DataCertPlatform/db_mysql"
//	"crypto/md5"
//	"encoding/hex"
//	_"大一下学期/github.com/go-sql-driver/mysql"
//)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}
//将用户信息保存到数据库中
func (u User) AddUser()(int64,error){
	//脱敏
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)//把脱敏的密码的md5值重新赋值给密码
	rs,err := db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",u.Phone,u.Password)
	//错误早发现早解决
	if err != nil{//保存数据遇到错误
		return -1,err
	}
	id, err := rs.RowsAffected()
	if err != nil{//保存数据遇到错误
		return -1,err
	}
	//id是影响到的行数
	return  id,nil
}
func (u User)QueryUser()(*User,error){
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)//把脱敏的密码的md5值重新赋值给密码
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",u.Phone,u.Password)//查询一条数据
	err := row.Scan(&u.Phone)//浏览，读取
	if err != nil{
		return nil,err
	}
	return &u,nil
}