package repository

import (
	"clean-arch-hicoll/pkg/domain"
	"database/sql"
)

type UploadRepository struct {
	db *sql.DB
}

func NewUploadRepository(db *sql.DB) *UploadRepository {
	return &UploadRepository{
		db: db,
	}
}

func (ur *UploadRepository) AddNewFileUpload(url string) error {
	sql := "INSERT INTO uploads (url) VALUES ($1)"

	_, err := ur.db.Exec(sql, url)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UploadRepository) GetAllFileUploads() ([]domain.Upload, error) {
	panic("not implemented") // TODO: Implement
}
