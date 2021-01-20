package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uint   `json:"-" gorm:"not null"`
	UserName string `json:"user_name" gorm:"-"`
	PostID   uint   `json:"post_id" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
}

func (c *Comment) Create(db *gorm.DB) (uint, error) {
	result := db.Create(c)
	return c.ID, result.Error
}

func (c *Comment) Select(db *gorm.DB, condition *map[string]interface{}, offset, limit int) ([]Comment, error) {
	comments := []Comment{}
	query := db.Offset(offset).Limit(limit)
	if condition != nil {
		query = query.Where(*condition)
	}
	records, err := query.Table("comments").Select("comments.id, comments.content, users.name").Joins("join users on comments.user_id = users.id").Rows()
	if err != nil {
		return comments, err
	}
	for records.Next() {
		var c Comment
		if err = records.Scan(&c.ID, &c.Content, &c.UserName); err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
