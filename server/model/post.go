package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PostForm struct {
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(128)"`
}

type Post struct {
	gorm.Model
	PostForm
}

func Insert(db *gorm.DB, value, imageUrl string) error {
	return db.Create(&Post{PostForm: PostForm{Name: value, ImageUrl: imageUrl}}).Error
}

func PostSelect(db *gorm.DB, cond *map[string]interface{}, offset, limit int) ([]Post, error) {
	posts := []Post{}
	query := db.Offset(offset).Limit(limit)
	if cond != nil {
		query = query.Where(*cond)
	}
	result := query.Find(&posts)
	return posts, result.Error
}
