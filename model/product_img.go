package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ImgPath   string `json:"img_path" gorm:"type:varchar(225) not null"`
	ProductId uint   `json:"product_id" gorm:"type:int(11) not null"`
}
