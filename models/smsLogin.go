package models

type SmsLogin struct {
	BizId string `form:"biz_id"`
	Phone string `form:"phone"`
	Code string	`form:"code"`
}