package controller

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestFetchHoges(t *testing.T) {
	e := echo.New()
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatalf("sqlmock err: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hoges` WHERE `hoges`.`deleted_at` IS NULL LIMIT 2 OFFSET 0")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "aaaaa").AddRow(2, "bbbbb"))

	req := httptest.NewRequest(http.MethodGet, "/hoges?offset=0&limit=2", nil)
	rec := httptest.NewRecorder()
	cx := model.DBContext{e.NewContext(req, rec), db}
	err = FetchHoges(&cx)
	if err != nil {
		t.Fatal(err)
	}
}
