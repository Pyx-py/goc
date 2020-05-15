package models

type User struct {
	BaseModel
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "admin_users"
}
