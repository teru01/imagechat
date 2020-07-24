package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func makePost(userID uint, name, image_url string) *model.Post {
	return &model.Post{
		UserID:   userID,
		Name:     name,
		ImageUrl: image_url,
	}
}

func createUser(t *testing.T, db *gorm.DB) {
	if err := test.CreateSeedData([]model.Creatable{
		&model.User{
			Name:     "myuser",
			Email:    "a@example.com",
			Password: "xxxxx",
		},
	}, db); err != nil {
		t.Fatal(err)
	}
}

func TestCanGetSpecificPostByPathParam(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)
	e := echo.New()

	createUser(t, db)
	target := makePost(1, "bar", "http://hoge.com/3434")
	if err := test.CreateSeedData([]model.Creatable{
		makePost(1, "hogehoge", "http://hoge.com/1212"),
		target,
		makePost(1, "woo", "http://hoge.com/5656"),
	}, db); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}
	cx.SetPath("/posts/:id")
	cx.SetParamNames("id")
	cx.SetParamValues("2")

	if err := FetchPost(&cx); err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	var response model.Post
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, target.ImageUrl, response.ImageUrl)
	assert.Equal(t, target.Name, response.Name)
	assert.Equal(t, "myuser", response.UserName)
}

func TestCanGetSpecificPostsByQueryParameter(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)
	e := echo.New()
	createUser(t, db)

	var posts []model.Creatable

	for i := 0; i < 2; i++ {
		posts = append(posts, makePost(1, fmt.Sprintf("name_%v", i), fmt.Sprintf("http://example.com/%v.png", i)))
	}
	targetPosts := []model.Creatable{
		makePost(1, "qwrty", "http://example.com/a.png"),
		makePost(1, "zxcvb", "http://example.com/b.png"),
	}
	posts = append(posts, targetPosts...)
	for i := 0; i < 2; i++ {
		posts = append(posts, makePost(1, fmt.Sprintf("name_%v", i), fmt.Sprintf("http://example.com/%v.png", i)))
	}
	if err := test.CreateSeedData(posts, db); err != nil {
		t.Fatal(err)
	}

	q := make(url.Values)
	q.Set("offset", "2")
	q.Set("limit", "2")
	req := httptest.NewRequest(http.MethodGet, "/posts?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}

	if err := FetchPosts(&cx); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	var response []model.Post
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, targetPosts[0].(*model.Post).Name, response[0].Name)
	assert.Equal(t, targetPosts[0].(*model.Post).ImageUrl, response[0].ImageUrl)
	assert.Equal(t, "myuser", response[0].UserName)
	assert.Equal(t, targetPosts[1].(*model.Post).Name, response[1].Name)
	assert.Equal(t, targetPosts[1].(*model.Post).ImageUrl, response[1].ImageUrl)
	assert.Equal(t, "myuser", response[1].UserName)
}
