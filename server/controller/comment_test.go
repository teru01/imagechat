package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func TestCanPostComment(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)

	e := echo.New()

	userID := createTestUser(t, db)
	post1 := model.NewPost(userID, "bar", "http://hoge.com/3434")
	post2 := model.NewPost(userID, "mypost", "http://hoge.com/3434")
	if err := test.CreateSeedData([]model.Creatable{
		post1,
		post2,
	}, db); err != nil {
		t.Fatal(err)
	}
	comment1 := model.NewComment(userID, post2.ID, "hello")
	comment2 := model.NewComment(userID, post2.ID, "yap")
	if err := test.CreateSeedData([]model.Creatable{
		model.NewComment(userID, post1.ID, "great!"),
		comment1,
		comment2,
	}, db); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}
	cx.SetPath("/:post_id/comments")
	cx.SetParamNames("post_id")
	cx.SetParamValues("2")

	if err := FetchComments(&cx); err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	var responses []model.Comment
	if err := json.Unmarshal(rec.Body.Bytes(), &responses); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(responses), 2)
	assert.Equal(t, comment1.Content, responses[0].Content)
	assert.Equal(t, comment2.Content, responses[1].Content)
}
