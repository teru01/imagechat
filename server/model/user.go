package model

import "github.com/jinzhu/gorm"

type UserForm struct {
	Name     string `json: "name" gorm:"type:varchar(255)"`
	Email    string `json: "email" gorm:"type:varchar(255)"`
	Password string `json: "password" gorm:"type:varchar(65500)"`
}

type User struct {
	gorm.Model
	UserForm
}

func Create(db *gorm.DB, userForm *UserForm) (UserForm, error) {
	if err := db.Create(&User{UserForm: *userForm}).Error; err != nil {
		return UserForm{}, err
	}
	return UserForm{Name: userForm.Name, Email: userForm.Email}, nil
}
