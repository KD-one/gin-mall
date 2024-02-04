package model

import "gorm.io/gorm"

// Notice 公告
type Notice struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(20) not null"`
	Content string `json:"content" gorm:"type:varchar(20) not null"`
}
