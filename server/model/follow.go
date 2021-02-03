package model

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	UserID       uint `json:"user_id" gorm:"not null"`
	FollowUserID uint `json:"follow_user_id" gorm:"not null"`
}

func (f *Follow) Create(db *gorm.DB) (uint, error) {
	result := db.Create(f)
	return f.ID, result.Error
}
