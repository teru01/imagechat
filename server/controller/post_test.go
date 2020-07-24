package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func TestCanGetSpecificPostByPathParam(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)
	e := echo.New()

	target := model.Post{
		Name:     "bar",
		ImageUrl: "http://hoge.com/3434",
	}
	if err := test.CreateSeedData([]model.Creatable{
		&model.Post{
			Name:     "hogehoge",
			ImageUrl: "http://hoge.com/1212",
		},
		&target,
		&model.Post{
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
}

func TestCanGetSpecificPostsByQueryParameter(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)
	e := echo.New()
	var posts []model.Creatable

	for i := 0; i < 2; i++ {
		posts = append(posts, &model.Post{
			Name:     fmt.Sprintf("name_%v", i),
			ImageUrl: fmt.Sprintf("http://example.com/%v.png", i),
		})
	}
	targetPosts := []model.Creatable{
		&model.Post{
			Name:     "qwrty",
			ImageUrl: "http://example.com/a.png",
		},
		&model.Post{
			Name:     "zxcvb",
			ImageUrl: "http://example.com/b.png",
		},
	}
	posts = append(posts, targetPosts...)
	for i := 0; i < 2; i++ {
		posts = append(posts, &model.Post{
			Name:     fmt.Sprintf("name_%v", i),
			ImageUrl: fmt.Sprintf("http://example.com/%v.png", i),
		})
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
	assert.Equal(t, targetPosts[1].(*model.Post).Name, response[1].Name)
	assert.Equal(t, targetPosts[1].(*model.Post).ImageUrl, response[1].ImageUrl)
}
