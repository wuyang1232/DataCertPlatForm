package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1、创世区块
	bc := blockchain.NewBlockChain()//封装
	fmt.Printf("创世区块哈希值：%x\n",bc.LastHash)
	//bc.SaveData([]byte("用户要保存到区块中的数据"))
	return
	//打开数据库
	db_mysql.Connect()
	//静态资源文件映射设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/img", "./static/img")
	beego.SetStaticPath("/css", "./static/css")
	//
	beego.Run()
}