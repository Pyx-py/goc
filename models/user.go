package models

type User struct {
	BaseModel
	UserName string `json:"username" gorm:"column:username" form:"userName"`
	Password string `json:"password" gorm:"column:password" form:"passWord"`
}

func (User) TableName() string {
	return "admin_users"
}
