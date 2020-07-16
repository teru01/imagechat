package model

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/teru01/image/server/database"
	"golang.org/x/crypto/bcrypt"
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
	return string(hashedPasswd), nil
}

func (user *User) Login(context *database.DBContext) (int, error) {
	var authenticatedUser User
	hashedPassword, err := hashPassword(user.Password)
	fmt.Println(hashedPassword)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	context.Db.Where("email = ?", user.Email).First(&authenticatedUser)
	if bcrypt.CompareHashAndPassword([]byte(authenticatedUser.Password), []byte(user.Password)) != nil {
		return http.StatusUnauthorized, fmt.Errorf("email or password is wrong")
	}
	sess, err := session.Get("session", context)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	sess.Options = &sessions.Options{
	  Path:     "/",
	  MaxAge:   86400 * 7,
	  HttpOnly: true,
	}
	sess.Values["user_id"] = user.Model.ID
	sess.Save(context.Request(), context.Response())
	return http.StatusOK, nil
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
