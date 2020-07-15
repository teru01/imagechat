package model

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PostForm struct {
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(128)"`
}

type Post struct {
	gorm.Model
	PostForm
}

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

func RegisterPost(c *model.DBContext) error {
	h := new(model.PostForm)

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


func Insert(db *gorm.DB, value, imageUrl string) error {
	return db.Create(&Post{PostForm: PostForm{Name: value, ImageUrl: imageUrl}}).Error
}

func SelectPost(db *gorm.DB, cond *map[string]interface{}, offset, limit int) ([]Post, error) {
	posts := []Post{}
	query := db.Offset(offset).Limit(limit)
	if cond != nil {
		query = query.Where(*cond)
	}
	result := query.Find(&posts)
	return posts, result.Error
}
