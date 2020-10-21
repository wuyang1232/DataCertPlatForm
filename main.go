package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock()
	block1 := blockchain.NewBlock(block0.Height+1, block0.Hash, []byte("a"))
	fmt.Println(block0,block1)
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