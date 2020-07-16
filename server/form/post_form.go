package form

type PostForm struct {
	Name     string `json:"name" gorm:"type:varchar(255)"`
	// ImageUrl string `json:"image_url" gorm:"type:varchar(128)"`
}
