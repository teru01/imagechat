package model

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"

	"cloud.google.com/go/storage"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Uploader interface {
	Upload(name string, file *multipart.File) (string, error)
}

type GCPUploader struct{}
type LocalUploader struct{}

func (uploader *GCPUploader) Upload(name string, file *multipart.File) (string, error) {
	fileExtension := path.Ext(name)
	imageUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", os.Getenv("BUCKET_NAME"), name)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return imageUrl, err
	}
	writer := client.Bucket(os.Getenv("BUCKET_NAME")).Object(name).NewWriter(ctx)
	defer writer.Close()
	writer.ContentType = convertExtensionToContentType(fileExtension)

	_, err = io.Copy(writer, *file)
	if err != nil {
		return imageUrl, err
	}
	return imageUrl, nil
}

func (uploader *LocalUploader) Upload(name string, file *multipart.File) (string, error) {
	imageUrl := fmt.Sprintf("http://localhost:8080/%s", name)
	fp, err := os.Create(path.Join(os.Getenv("IMG_ROOT"), name))
	if err != nil {
		return imageUrl, err
	}
	defer fp.Close()
	_, err = io.Copy(fp, *file)
	if err != nil {
		return imageUrl, err
	}
	return imageUrl, nil
}

func convertExtensionToContentType(ext string) string {
	return "image/" + ext
}
