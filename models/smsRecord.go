package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	Message string
	TimeStamp int64
}

//根据用户提交的手机号和短信验证码查询验证码是否正确及正常
func QuerySmsRecord(bizId string,phone string,code string)(*SmsRecord,error){
	var sms SmsRecord
	row := db_mysql.Db.QueryRow("select  biz_id, timestamp from sms_record where biz_id = ? and phone = ? andcode = ?",bizId,phone,code)
	err := row.Scan(&sms.BizId,&sms.TimeStamp,)
	if err != nil{
		return nil,err
	}
	return &sms,nil
}

//项数据库当中插入验证码记录，该记录由阿里云第三方返回
func (s SmsRecord) SaveSmsRecord()(int64, error){
	rs,err := db_mysql.Db.Exec("insert into sms_record(biz_id, phone, code, status, message, timestamp) values(?,?,?,?,?,?)",
		s.BizId, s.Phone,s.Code,s.Status,s.Message,s.TimeStamp)
	if err != nil{
		return -1, err
	}
	return rs.RowsAffected()
}