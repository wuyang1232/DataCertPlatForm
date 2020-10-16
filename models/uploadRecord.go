package models

import "DataCertPlatform/db_mysql"

type UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int64
	FileCert string
	FileTitle string
	CertTime int64
}
//把一条认证数据保存到数据库中
func (u UploadRecord) SaveRecord()(int64,error){
	re, err := db_mysql.Db.Exec("insert into upload_record(user_id, file_name, file_size, file_cert, file_title, cert_time)" +
		"values(?,?,?,?,?,?)",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileCert,u.CertTime)
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

}
