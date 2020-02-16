package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key;auto_increment;not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UserName  string    `json:"user_name" gorm:"-"`
	UserEmail string    `json:"user_email" gorm:"-"`
}

func CreateComment(db *gorm.DB, comment *Comment) (*Comment, error) {
	comment.CreatedAt = time.Now()
	return comment, db.Create(comment).Error
}

func FetchComments(db *gorm.DB, condition *map[string]interface{}) ([]Comment, error) {
	comments := []Comment{}
	if condition != nil {
		db = db.Where(*condition)
	}
	records, err := db.Table("comments").Select("comments.id, comments.content, users.name").Joins("join users on comments.user_id = users.id").Rows()
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
