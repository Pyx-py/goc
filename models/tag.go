package models

type Tag struct {
	BaseModel
	TagName      string `json:"tag_name" gorm:"column:tag_name"`
	CreateUser   string `json:"create_user" gorm:"create_user"`
	ModifiedUser string `json:"modified_user" gorm:"modified_user"`
	Status       uint   `json:"status" gorm:"status"`
}

func (Tag) TableName() string {
	return "tags"
}
