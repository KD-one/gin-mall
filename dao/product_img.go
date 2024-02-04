package dao

import "gin-mall/model"

func CreateProductImg(productImg *model.ProductImg) error {
	return DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

// GetProductImgsById 根据商品id查询商品的所有图片
func GetProductImgsById(id uint) (productImgs []*model.ProductImg, err error) {
	err = DB.Where("product_id = ?", id).Find(&productImgs).Error
	return productImgs, err
}
