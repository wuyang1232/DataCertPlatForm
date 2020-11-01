package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	//先准备一条区块链
	blockchain.NewBlockChain()
	//1、创世区块
	//bc := blockchain.NewBlockChain() //封装
	//bc.SaveData([]byte("小憨憨"))
	//blocks,err := bc.QueryAllBlocks()
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	////blocks是一个切片
	//for _, block := range blocks{
	//	fmt.Printf("区块高度：%d,区块的hash:%x,PrevHash:%x,区块内数据：%s\n",block.Height,block.Hash,block.PrevHash,string(block.Data))
	//
	//}
	//return
	//打开数据库
	db_mysql.Connect()
	//静态资源文件映射设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/img", "./static/img")
	beego.SetStaticPath("/css", "./static/css")
	//
	beego.Run()
}
