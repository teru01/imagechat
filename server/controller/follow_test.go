package controller

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	"github.com/teru01/image/server/test"
)

func TestFollow(t *testing.T) {
	db := test.SetUpDB()
	defer test.TearDownDB(db)
	e := echo.New()

	folowee := createTestUser(t, db)
	folower := createTestUser(t, db)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	cx := database.DBContext{e.NewContext(req, rec), db}

	store := sessions.NewCookieStore([]byte("secret-key"))
	cx.Set("_session_store", store)
	if _, err := model.NewSession(&folowee, &cx); err != nil {
		t.Fatal(err)
	}
	cx.SetPath("/follow/:user_id")
	cx.SetParamNames("user_id")
	cx.SetParamValues(strconv.Itoa(int(folower.ID)))
	if err := Follow(&cx); err != nil {
		t.Fatal(err)
	}
	follow := model.Follow{}
	db.First(&follow)
	assert.Equal(t, follow.UserID, folowee.ID)
	assert.Equal(t, follow.FollowUserID, folower.ID)
}
