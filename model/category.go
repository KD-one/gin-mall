package model

import "gorm.io/gorm"

// Category 商品分类
type Category struct {
	gorm.Model
	Name string
	//ParentId uint
	//Children []Category
}
