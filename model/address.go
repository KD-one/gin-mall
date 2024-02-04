package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId  uint   `json:"user_id" gorm:"not null"`
	Name    string `json:"name" gorm:"type:varchar(20) not null"`
	Phone   string `json:"phone" gorm:"type:varchar(11) not null"`
	Address string `json:"address" gorm:"type:varchar(255) not null"`
}
