package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
)

var db *gorm.DB

func TestFetchPost(t *testing.T) {
	setUp()
	defer tearDown()
	e := echo.New()

	target := model.Post{
		Name:     "bar",
		ImageUrl: "http://hoge.com/3434",
	}
	targetJson := `{"name":"bar", "image_url":"http://hoge.com/3434"}`
	if err := createSeedData([]model.Post{
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

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}
	cx.SetPath("/posts/:id")
	cx.SetParamNames("id")
	cx.SetParamValues("2")

	err := FetchPost(&cx) //routeで/posts/:idのパースが行われないから
	if err != nil {
		t.Error(err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status code. Expected %v got %v", http.StatusOK, rec.Code)
	}
	var response model.Post
	if err = json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.ImageUrl != target.ImageUrl || response.Name != target.Name {
		t.Errorf("unexpected response. Expected %v got %v", targetJson, response)
	}
}

func createSeedData(items []model.Post, db *gorm.DB) error {
	for _, i := range items {
		if err := db.Create(&i).Error; err != nil {
			return err
		}
	}
	return nil
}

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

func setUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	db = database.ConnectDB(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "localhost", os.Getenv("MYSQL_TEST_DATABASE"))
	db.LogMode(true)
	InitializeDB(db)
}

func tearDown() {
	ResetDB(db)
	db.Close()
}
