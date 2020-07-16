package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/teru01/image/server/database"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Comments []Comment
}

func userAvailable(db *gorm.DB, name, email string) (bool, string) {
	var users []User
	db.Where("name = ?", name).Find(&users)
	if len(users) != 0 {
		return false, fmt.Sprintf("name %v is already used", name)
	}
	db.Where("email = ?", email).Find(&users)
	if len(users) != 0 {
		return false, fmt.Sprintf("email %v is already used", email)
	}
	return true, ""
}

func (user *User) SignUp(db *gorm.DB) error {
	ok, msg := userAvailable(db, user.Name, user.Email)
	if !ok {
		return fmt.Errorf(msg)
	}

	hashed, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

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

func (user *User) Login(context *database.DBContext) error {
	var authenticatedUser User
	context.Db.Where("email = ? AND password = ?", user.Name, user.Password).First(&authenticatedUser)
	sess, err := session.Get("session", context)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
	  Path:     "/",
	  MaxAge:   86400 * 7,
	  HttpOnly: true,
	}
	sess.Values["user_id"] = authenticatedUser.Model.ID
	sess.Save(context.Request(), context.Response())
	return nil
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
