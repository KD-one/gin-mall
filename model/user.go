package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(20) not null"`
	Password string `json:"password" gorm:"type:varchar(255) not null"`
	Email    string `json:"email" gorm:"type:varchar(255) not null"`
	NickName string `json:"nickname" gorm:"type:varchar(20) not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255) not null"`
	Status   int    `json:"status" gorm:"type:int(11) not null"` // 1:正常 2:禁用
	Money    string `json:"money" gorm:"type:varchar(255) not null"`
}
