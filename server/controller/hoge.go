package controller

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
)

func uploadImage(fileHeader *multipart.FileHeader, writer io.Writer) error {
	srcImg, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer srcImg.Close()

	_, err = io.Copy(writer, srcImg)
	if err != nil {
		return err
	}
	return nil
}

func RegisterHoge(c *model.DBContext) error {
	h := new(model.HogeForm)

	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	name := uuid.New().String() + path.Ext(fileHeader.Filename)
	var imageUrl string
	// GCS処理
	if os.Getenv("ENV_TYPE") == "prod" {
		imageUrl = fmt.Sprintf("https://storage.googleapis.com/%s/%s", os.Getenv("BUCKET_NAME"), name)
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return err
		}
		writer := client.Bucket(os.Getenv("BUCKET_NAME")).Object(name).NewWriter(ctx)
		defer writer.Close()
		writer.ContentType = c.Request().Header.Get("Content-Type")
		if err = uploadImage(fileHeader, writer); err != nil {
			return err
		}
	} else {
		imageUrl = fmt.Sprintf("http://localhost:8080/%s", name)
		fp, err := os.Create(path.Join(os.Getenv("IMG_ROOT"), name))
		if err != nil {
			return err
		}
		defer fp.Close()
		if err = uploadImage(fileHeader, fp); err != nil {
			return err
		}
	}

	if err := c.Bind(h); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := model.Insert(c.Db, h.Name, imageUrl); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func FetchHoges(c *model.DBContext) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	hoges, err := model.HogeSelect(c.Db, nil, offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hoges)
}

func FetchHoge(c *model.DBContext) error {
	hoges, err := model.HogeSelect(c.Db, &map[string]interface{}{"id": c.Param("id")}, 0, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(hoges) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, hoges[0])
}
