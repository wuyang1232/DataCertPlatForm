package models

type UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int
	FileCert string
	FileTitle string
	CertTime int
}
