package form

type PostForm struct {
	Name     string `json:"name" form:"name" gorm:"type:varchar(255)"`
}
