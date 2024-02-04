package model

import "gorm.io/gorm"

// Favorite 收藏夹
type Favorite struct {
	gorm.Model
	User      User    `json:"user" gorm:"not null ForeignKey:UserId"`
	UserId    uint    `json:"user_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"not null ForeignKey:ProductId"`
	ProductId uint    `json:"product_id" gorm:"not null"`
	Boos      User    `json:"boos" gorm:"not null ForeignKey:BoosId"`
	BoosId    uint    `json:"boos_id" gorm:"not null"`
}
