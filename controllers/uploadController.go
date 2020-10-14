package controllers

import (
	"beego"
	"fmt"
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
	file, header, err := u.GetFile("file")
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
	//isJpg := strings.HasSuffix()

	//文件大小200kb
	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")
	if header.Size / 1024 > fileSize{
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传合适格式的文件")
		return
	}
	fmt.Println("文件的大小:",header.Size)//字节大小

	fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("已获取到上传文件")

	//u.TplName = "login.html"
}
