package dao

import "gin-mall/model"

// GetCarousels 获取所有轮播图
func GetCarousels() (carousels []*model.Carousel, err error) {
	err = DB.Model(&model.Carousel{}).Find(&carousels).Error
	return carousels, err
}
