package model

import "gorm.io/gorm"

// Carousel 轮播图
type Carousel struct {
	gorm.Model
	ImgPath   string `json:"img_path" gorm:"type:varchar(255) not null"`
	ProductId uint   `json:"product_id" gorm:"not null"`
}
