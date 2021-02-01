package test

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
)

var m *migrate.Migrate

func InitializeDB(db *gorm.DB) {
	var err error
	driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})
	m, err = migrate.NewWithDatabaseInstance("file://../../mysql/migrations", "myapp", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func ResetDB(db *gorm.DB) {
	if err := m.Drop(); err != nil {
		log.Fatal(err)
	}
}

func SetUpDB() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil && os.Getenv("CIRCLECI") != "" {
		log.Fatal(err)
	}
	db := database.ConnectDB(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_TEST_HOST"), os.Getenv("MYSQL_TEST_DATABASE"), "0")
	InitializeDB(db)
	return db
}

func TearDownDB(db *gorm.DB) {
	ResetDB(db)
	db.Close()
}

func CreateSeedData(items []model.Creatable, db *gorm.DB) error {
	for _, i := range items {
		if _, err := i.Create(db); err != nil {
			return err
		}
	}
	return nil
}
