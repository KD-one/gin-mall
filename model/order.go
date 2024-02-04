package model

import "gorm.io/gorm"

// Order 订单
type Order struct {
	gorm.Model
	UserId    uint   `json:"user_id" gorm:"type:int(11) not null"`
	ProductId uint   `json:"product_id" gorm:"type:int(11) not null"`
	BoosId    uint   `json:"boos_id" gorm:"type:int(11) not null"`
	AddressId uint   `json:"address_id" gorm:"type:int(11) not null"`
	Type      uint   `json:"type" gorm:"type:int(11) not null"` // 1.已支付 2.未支付
	Num       int    `json:"num" gorm:"type:int(11) not null"`
	OrderNum  uint   `json:"order_num" gorm:"type:int(11) not null"`
	Price     string `json:"price" gorm:"type:varchar(20) not null"`
}
