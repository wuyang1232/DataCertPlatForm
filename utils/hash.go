package utils

import (
	//"DataCertPlatform/blockchain"
	//"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串数据进行MD5hash计算
func Md5HashString(data string)string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

//读取io流中的数据并对数据进行hash计算,返回sha256，hash值
func Sha256HashReader(reader io.Reader)(string,error){
	sha256Hash := sha256.New()
	readerBytes,err := ioutil.ReadAll(reader)
	if err != nil{
		return "",err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

func Md5HashReader(reader io.Reader)(string,error){
	md5Hash := md5.New()
	readerBytes, err := ioutil.ReadAll(reader)
	//fmt.Println("读取到的文件",readerBytes)
	if err != nil{
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

//对区块数据进行sha256hash计算
func SHA256HashBlock(bs []byte)[]byte{
	//将block结构体数据转换为[]byte切片类型
	//heightBytes,_ := Int64ToByte(block.Height)
	//timeStampBytes,_ := Int64ToByte(block.TimeStamp)
	//versionBytes := StringToBytes(block.Version)
	//
	//var blockBytes []byte
	////bytes.Join拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//},[]byte{})

	//将转换后的[]byte字节切片输入write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return hash
}

