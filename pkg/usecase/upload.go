package usecase

import (
	"clean-arch-hicoll/pkg/domain"
	"mime/multipart"
)

type UploadUsecase struct {
	ur domain.UploadRepository
	us domain.UploadStorage
}

func NewUploadUsecase(ur domain.UploadRepository, us domain.UploadStorage) *UploadUsecase {
	return &UploadUsecase{
		ur: ur,
		us: us,
	}
}

func (uu *UploadUsecase) UploadNewFile(file multipart.File, fileName string) error {
	url, err := uu.us.UploadNewFile(file, fileName)
	if err != nil {
		return err
	}

	err = uu.ur.AddNewFileUpload(url)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UploadUsecase) GetAllFileUploads() ([]domain.Upload, error) {
	panic("not implemented") // TODO: Implement
}
