package model

import (
	"fmt"
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
	users := []User{}
	if condition != nil {
		db = db.Where(*condition)
	}
	records := db.Preload("Comments").Find(&users)
	fmt.Println(comments)
	fmt.Println(users)
	return comments, records.Error
}
