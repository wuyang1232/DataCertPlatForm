package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock()//创建创世区块
	//fmt.Println(block0)

	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		nil)
	fmt.Printf("block0的hash：%x\n",block0.Hash)
	fmt.Printf("block1的哈希:%x\n",block1.Hash)
	fmt.Printf("block1的PrevHash:%x\n",block1.PrevHash)

	//序列化----将数据从内存形式转换成可以持久化存储在硬盘上或网络上传输的形式
	//反序列化----将数据从文件或者网络中读取，
	//只有序列化以后的数据才可以进行传输
	blockJson,_ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block：",string(blockJson))
	blockXml,_ := xml.Marshal(block0)
	fmt.Println("通过json序列化以后的block:",string(blockXml))
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