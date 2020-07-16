package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/teru01/image/server/form"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Comments []Comment
}

func (user *User) CreateUser(db *gorm.DB, userForm form.UserForm) error {
	hashed, err := hashPassword(userForm.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	user.Name = userForm.Name
	user.Email = userForm.Email

	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func hashPassword(original string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(original), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hashedPasswd), nil
}

// func UpdateUser(db *gorm.DB, user *User, m map[string]interface{}) (*User, error) {
// 	if err := db.Model(user).Update(m).Error; err != nil {
// 		return nil, err
// 	}
// 	return &User{ID: user.ID, Name: user.Name}, nil
// }

// func DeleteUser(db *gorm.DB, user *User) error {
// 	return db.Delete(user).Error
// }
