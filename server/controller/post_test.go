package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/proullon/ramsql/driver"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"

)

func getDBMock() (*gorm.DB, error) {
	gdb, err := gorm.Open("ramsql", "TestDB")
	if err != nil {
		return nil, err
	}
	return gdb, nil
}

func TestFetchPost(t *testing.T) {
	e := echo.New()
	db, err := getDBMock()
	if err != nil {
		t.Fatalf("db open err: %v", err)
	}
	defer db.Close()

	// mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`deleted_at` IS NULL LIMIT 2 OFFSET 0")).
		// WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "aaaaa").AddRow(2, "bbbbb"))

	InitializeDB(db)

	target := model.Post{
		Name:     "bar",
		ImageUrl: "http://hoge.com/3434",
	}
	targetJson := `{"name":"bar", "image_url":"http://hoge.com/3434"}`
	if err := createSeedData([]model.Post {
		model.Post{
			Name:     "hogehoge",
			ImageUrl: "http://hoge.com/1212",
		},
		target,
		model.Post{
			Name:     "foor",
			ImageUrl: "http://hoge.com/5656",
		},
	}, db); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/posts/2", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}
	err = FetchPosts(&cx)
	if err != nil {
		t.Error(err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status code. Expected %v got %v", http.StatusOK, rec.Code)
	}
	response := rec.Body.String()
	if response != targetJson {
		t.Errorf("unexpected response. Expected %v got %v", targetJson, response)
	}
}

func createSeedData(items []model.Post, db *gorm.DB) error {
	for i := range items {
		if err := db.Create(&i).Error; err != nil {
			return err
		}
	}
	return nil
}

func InitializeDB(db *gorm.DB) {
	// db.Exec(`CREATE TABLE posts (
	// 	id serial primary key,
	// 	created_at timestamp,
	// 	updated_at timestamp,
	// 	deleted_at timestamp,
	// 	name varchar(255),
	// 	image_url varchar(255))`)
		// db.Exec(`CREATE TABLE posts (id BIGSERIAL PRIMARY KEY, street TEXT, street_number INT);`)
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Comment{})
}
