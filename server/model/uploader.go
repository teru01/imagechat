package model

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/teru01/image/server/form"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Uploader interface {
	Upload(name string, file *multipart.File) (string, error)
}

type GCPUploader struct {}
type LocalUploader struct {}

func (uploader *GCPUploader) Upload(name string, file *multipart.File) (string, error) {
	var imageUrl string
	fileExtension := path.Ext(name)
	imageUrl = fmt.Sprintf("https://storage.googleapis.com/%s/%s", os.Getenv("BUCKET_NAME"), name)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return imageUrl, err
	}
	writer := client.Bucket(os.Getenv("BUCKET_NAME")).Object(name).NewWriter(ctx)
	defer writer.Close()
	writer.ContentType = convertExtensionToContentType(fileExtension)
	if err = uploadImage(fileHeader, writer); err != nil {
		return imageUrl, err
	}
	return imageUrl, nil
}

func (uploader *LocalUploader) Upload(name string, file *multipart.File) (string, error) {
	imageUrl = fmt.Sprintf("http://localhost:8080/%s", name)
	fp, err := os.Create(path.Join(os.Getenv("IMG_ROOT"), name))
	if err != nil {
		return err
	}
	defer fp.Close()
	if err = uploadImage(fileHeader, fp); err != nil {
		return err
	}
}

func uploadImage(fileHeader *multipart.FileHeader, writer io.Writer) error {
	srcImg, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer srcImg.Close()

	_, err = io.Copy(writer, srcImg)
	if err != nil {
		return err
	}
	return nil
}

func convertExtensionToContentType(ext string) string {
	return "image/" + ext
}
