package model

import "gorm.io/gorm"

// Cart 购物车
type Cart struct {
	gorm.Model
	UserId    uint `json:"user_id" gorm:"not null"`
	ProductId uint `json:"product_id" gorm:"not null"`
	BoosId    uint `json:"boos_id" gorm:"not null"` // 商家id
	Num       uint `json:"num" gorm:"not null"`
	Price     uint `json:"price" gorm:"not null"`
	Total     uint `json:"total" gorm:"not null"`
	IsDelete  bool `json:"is_delete" gorm:"not null"`
}
