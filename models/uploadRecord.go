package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
)

type UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int64
	FileCert string
	FileTitle string
	CertTime int64
	CertTimeFormat string
}
//把一条认证数据保存到数据库中
func (u UploadRecord) SaveRecord()(int64,error){
	re, err := db_mysql.Db.Exec("insert into upload_record(user_id, file_name, file_size, file_cert, file_title, cert_time)" +
		"values(?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
	if err != nil{
		return -1,err
	}
	id, err := re.RowsAffected()
	if err != nil{
		return -1,err
	}
	return id,nil
}
//根据用户id查询符合条件的认证数据记录
func QueryRecordsByUserId(userId int)([]UploadRecord,error){
	rs,err := db_mysql.Db.Query("select id, user_id, file_name, file_size, file_cert, file_title, cert_time from upload_record where user_id = ?",userId)
	if err != nil{
		return nil,err
	}
	records := make([]UploadRecord, 0)//容器
	for rs.Next(){
		var record UploadRecord
		err := rs.Scan(&record.Id,&record.UserId,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime)
		if err != nil{
			return nil, err
		}
		//整形转换成字符串
		tStr := utils.TimeFormat(record.CertTime,utils.TIME_FORMAT_ONE)
		record.CertTimeFormat = tStr
		//utils.TimeFormat(record.CertTime,utils.TIME_FORMAT_ONE)
		records = append(records,record)
	}
	return records,nil
}
