package domain

import "mime/multipart"

type Upload struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type UploadRepository interface {
	AddNewFileUpload(url string) error
	GetAllFileUploads() ([]Upload, error)
}

type UploadUsecase interface {
	UploadNewFile(file multipart.File, fileName string) error
	GetAllFileUploads() ([]Upload, error)
}

type UploadStorage interface {
	UploadNewFile(file multipart.File, fileName string) (string, error)
}
