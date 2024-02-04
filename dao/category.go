package dao

import "gin-mall/model"

// GetCategorys 获取所有轮播图
func GetCategorys() (categories []*model.Category, err error) {
	err = DB.Model(&model.Category{}).Find(&categories).Error
	return categories, err
}
