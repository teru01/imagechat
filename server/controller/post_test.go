package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func TestFetchPost(t *testing.T) {
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

	err := FetchPost(&cx) //routeで/posts/:idのパースが行われないから
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, rec.Code, http.StatusOK)
	var response model.Post
	if err = json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, response.ImageUrl, target.ImageUrl)
	assert.Equal(t, response.Name, target.Name)
}
