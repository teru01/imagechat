package model

import (
	"mime/multipart"
	"path"
	"strings"

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

func (p *Post) Submit(db *gorm.DB, fileHeader *multipart.FileHeader, postForm form.PostForm, uploader Uploader) error {
	fileExtension := strings.ToLower(path.Ext(fileHeader.Filename))
	name := uuid.New().String() + fileExtension

	srcImg, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer srcImg.Close()

	imageUrl, err := uploader.Upload(name, &srcImg)
	if err != nil {
		return err
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
