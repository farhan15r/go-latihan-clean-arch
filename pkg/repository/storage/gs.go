package storage

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type UploadStorage struct {
	client     *storage.Client
	bucketName string
	projectID  string
	uploadPath string
}

func NewUploadStorage() *UploadStorage {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "config/credentials.json")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	return &UploadStorage{
		client:     client,
		bucketName: "bucket-latihan-farhan",
		projectID:  "latihan-vertex",
		uploadPath: "uploads/",
	}
}

func (us *UploadStorage) UploadNewFile(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	obj := us.client.Bucket(us.bucketName).Object(us.uploadPath + fileName)
	obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader)
	wc := obj.NewWriter(ctx)

	_, err := io.Copy(wc, file)
	if err != nil {
		return "", err
	}

	err = wc.Close()
	if err != nil {
		return "", err
	}

	URL := "https://storage.googleapis.com/" + us.bucketName + "/" + us.uploadPath + fileName

	return URL, nil
}
