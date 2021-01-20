package model

import (
	"mime/multipart"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/form"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"not null"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(128)"`
	UserName string `json:"user_name" gorm:"-"`
}

func (p *Post) Submit(c *database.DBContext, fileHeader *multipart.FileHeader, postForm form.PostForm, uploader Uploader) error {
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
	p.ImageUrl = imageUrl
	p.Name = postForm.Name
	p.UserID = GetAuthSessionData(c, "user_id").(uint)
	_, err = p.Create(c.Db)
	return err
}

func (p *Post) Create(db *gorm.DB) (uint, error) {
	result := db.Create(p)
	return p.ID, result.Error
}

func (p *Post) Select(db *gorm.DB, condition *map[string]interface{}, offset, limit int) ([]Post, error) {
	posts := []Post{}
	query := db.Offset(offset).Limit(limit)
	if condition != nil {
		query = query.Where(*condition)
	}
	records, err := query.Table("posts").Select("posts.id, posts.name, posts.image_url, users.name").Joins("left join users on posts.user_id = users.id").Rows()
	if err != nil {
		return posts, err
	}
	for records.Next() {
		var p Post
		if err = records.Scan(&p.ID, &p.Name, &p.ImageUrl, &p.UserName); err != nil {
			return posts, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
