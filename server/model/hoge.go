package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type HogeForm struct {
	Name string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(65536)"`
}

type Hoge struct {
	gorm.Model
	HogeForm
}

func Insert(db *gorm.DB, value ,imageUrl string) error {
	return db.Create(&Hoge{HogeForm: HogeForm{Name: value, ImageUrl: imageUrl}}).Error
}

func HogeSelect(db *gorm.DB, cond *map[string]interface{} ,offset , limit int) ([]Hoge, error) {
	hoges := []Hoge{}
	query := db.Offset(offset).Limit(limit)
	if cond != nil {
		query = query.Where(*cond)
	}
	result := query.Find(&hoges)
	return hoges, result.Error
}
