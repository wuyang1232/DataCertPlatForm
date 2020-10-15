package controllers

import (
	"beego"
	"fmt"
	"os"
	"strings"
)

//该控制器结构体用于处理文件上传功能
type UploadController struct {
	beego.Controller
}
//该post方法用于处理该用户在客户端提交的文件认证
func (u *UploadController)Post(){
	//1、解析用户上传的数据及文件信息
	//用户上传的自定义标题
	title := u.Ctx.Request.PostFormValue("index_title")//用户输入的标题

	//用户上传的文件
	file, header, err := u.GetFile("file")//封装好了。下面可以通过该名字获取文件
	defer file.Close()

	if err != nil{//解析客户端的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	fmt.Println("自定义的标题:",title)
	//获得到了上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	if fileType != "jpg" && fileType != "png"{
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传合适格式的文件")
		return
	}
	//isJpg := strings.HasSuffix(header.Filename,".jpg")

	//文件大小200kb
	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")
	if header.Size / 1024 > fileSize{
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传合适格式的文件")
		return
	}
	fmt.Println("文件的大小:",header.Size)//字节大小

	//perm:permission权限
	//权限的组成：a+b+c
	//	a: 文件所有者对文件的操作权限： 读4，写2，执行1
	//	b: 文件所有者所在组的用户操作权限： 读4，写2，执行1
	//	c: 其他用户的操作权限： 读4，写2，执行1

	//eg: m文件，权限651。什么意思？文件所有者6 = 4 + 2文件所有者有读写权限
	// 文件所有者所在组的用户有写权限(错)
	saveDir := "static/upload"
	//打开文件
	f, err := os.Open(saveDir)
	if err != nil{
		//创建文件夹
		//err = os.Mkdir(saveDir,777)
		//if err != nil{//打开文件夹失败，自己创建
			//创建文件夹
			err = os.Mkdir(saveDir,777)
			if err != nil{
				fmt.Println(err.Error())
				u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试！")
				return
			}

			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试！")
			return
		//}

		fmt.Println(err.Error())
		u.Ctx.WriteString("打开文件目录失败")
		return
		//打开目录遇到错误
	}
	fmt.Println("打开的文件夹：",f.Name())

	//文件名：文件路径 + 文件名 + “.” + 文件拓展名
	saveName :=saveDir +  "/" + header.Filename
	fmt.Println("要保存的文件名",saveName)
	//fromFile:文件
	//toFile:要保存的文件路径
	err = u.SaveToFile("file",saveName)//通过封装好的名字获取文件
	if err != nil{
		fmt.Println(err.Error())
		u.Ctx.WriteString("文件认证失败，请重试")
		return
	}

	fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("已获取到上传文件")

	//u.TplName = "login.html"
}
