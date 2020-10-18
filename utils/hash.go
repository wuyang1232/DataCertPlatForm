package utils

import (
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
	if err != nil{
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}
