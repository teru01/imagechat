package model

import "github.com/jinzhu/gorm"

type Creatable interface {
	Create(*gorm.DB) error
}
