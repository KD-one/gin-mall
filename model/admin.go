package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(20) not null"`
	Password string `json:"password" gorm:"type:varchar(255) not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255) not null"`
	// Role     string `json:"role" gorm:"type:varchar(20) not null"`
}
