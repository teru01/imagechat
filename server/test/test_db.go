package test

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
)

func InitializeDB(db *gorm.DB) {
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Comment{})
}

func ResetDB(db *gorm.DB) {
	db.DropTable(&model.Post{})
	db.DropTable(&model.User{})
	db.DropTable(&model.Comment{})
}

func SetUpDB() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	db := database.ConnectDB(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "localhost", os.Getenv("MYSQL_TEST_DATABASE"), "0")
	InitializeDB(db)
	return db
}

func TearDownDB(db *gorm.DB) {
	ResetDB(db)
	db.Close()
}

func CreateSeedData(items []model.Creatable, db *gorm.DB) error {
	for _, i := range items {
		if err := i.Create(db); err != nil {
			return err
		}
	}
	return nil
}
