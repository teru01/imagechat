package controller

import (
	"fmt"
	"regexp"
	"testing"
	

	"github.com/labstack/echo/v4"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/teru01/image/server/model"
)

var e *echo.echo

func setUp() {
	e := echo.New()

}
func TestFetchHoges(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock err: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM hoges OFFSET ? LIMIT ?`)
		.WithArgs(0)
		.WithArgs(2)
		.WillReturnRows(NewRows([]string{"id", "name"})
			.AddRow(1, "aaaaa")
			.AddRow(2, "bbbbb")
		)
	)
	mock.ExpectCommit()

	cx := model.DBContext{e.NewContext(), db}
	err := FetchHoges(cx)
}
