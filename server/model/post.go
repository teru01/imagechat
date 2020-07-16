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

type Post struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(128)"`
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

func (p *Post) Submit(db *gorm.DB, fileHeader *multipart.FileHeader, postForm form.PostForm) error {
	fileExtension := strings.ToLower(path.Ext(fileHeader.Filename))
	name := uuid.New().String() + fileExtension
	var imageUrl string

	if os.Getenv("ENV_TYPE") == "prod" {
		// GCS処理
		imageUrl = fmt.Sprintf("https://storage.googleapis.com/%s/%s", os.Getenv("BUCKET_NAME"), name)
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return err
		}
		writer := client.Bucket(os.Getenv("BUCKET_NAME")).Object(name).NewWriter(ctx)
		defer writer.Close()
		writer.ContentType = convertExtensionToContentType(fileExtension)
		if err = uploadImage(fileHeader, writer); err != nil {
			return err
		}
	} else {
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

	return p.Insert(db, postForm.Name, imageUrl)
}

func (p *Post) Insert(db *gorm.DB, value, imageUrl string) error {
	return db.Create(&Post{
		Name:     value,
		ImageUrl: imageUrl,
	}).Error
}

func (p *Post) Select(db *gorm.DB, condition *map[string]interface{}, offset, limit int) ([]Post, error) {
	posts := []Post{}
	query := db.Offset(offset).Limit(limit)
	if condition != nil {
		query = query.Where(*condition)
	}
	result := query.Find(&posts)
	return posts, result.Error
}
