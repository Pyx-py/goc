package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"column:id"`
	CreatedAt time.Time  `json:"title" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updateAt" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}

type Article struct {
	BaseModel
	Title        string `json:"title" gorm:"column:title"`
	Content      string `json:"content" gorm:"column:content"`
	ModifiedUser string `json:"modified_user" gorm:"column:modified_user"`
	Like         uint   `json:"like" gorm:"column:like"`
}

func (Article) TableName() string {
	return "articles"
}
