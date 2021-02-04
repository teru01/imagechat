package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func createTestUser(t *testing.T, db *gorm.DB) model.User {
	user := model.User{
		Name:     "myuser",
		Email:    "a@example.com",
		Password: "xxxxx",
	}
	if err := test.CreateSeedData([]model.Creatable{
		&user,
	}, db); err != nil {
		t.Fatal(err)
	}
	return user
}

func TestCreateUser(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)

	e := echo.New()
	user_1 := model.NewUser("hoge", "foo@example.com", "aaaa1111")
	user_2 := model.NewUser("bar", "baro@example.com", "aaaa1111")
	user_3 := model.NewUser("qwe", "qwe@example.com", "aaaa1111")
	if err := test.CreateSeedData([]model.Creatable{
		user_1,
		user_2,
		user_3,
	}, db); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}
	cx.SetPath("/users/:id")
	cx.SetParamNames("id")
	cx.SetParamValues(strconv.Itoa(int(user_2.ID)))

	if err := FetchUser(&cx); err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	var response model.User
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user_2.Name, response.Name)
}
