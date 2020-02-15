package model

import "github.com/jinzhu/gorm"

type User struct {
	ID       int    `json: "id" gorm:"type:int;primary_key;auto_increment"`
	Name     string `json: "name" gorm:"type:varchar(255)"`
	Email    string `json: "email" gorm:"type:varchar(255)"`
	Password string `json: "password" gorm:"type:varchar(65500)"`
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return &User{ID: user.ID, Name: user.Name}, nil
}

func UpdateUser(db *gorm.DB, user *User, m map[string]interface{}) (*User, error) {
	if err := db.Model(user).Update(m).Error; err != nil {
		return nil, err
	}
	return &User{ID: user.ID, Name: user.Name}, nil
}
