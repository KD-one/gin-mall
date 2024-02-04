package dao

import (
	"gin-mall/model"
)

func CreateProduct(product *model.Product) error {
	return DB.Model(&model.Product{}).Create(&product).Error
}

// GetProductCountByCondition 根据condition条件查询商品总数
func GetProductCountByCondition(condition map[string]interface{}) (total int64, err error) {
	err = DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return total, err
}

// GetProductsByCondition 根据condition条件分页查询商品
func GetProductsByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return products, err
}

// GetProductCountByName 根据商品名查询商品总数
func GetProductCountByName(name string) (total int64, err error) {
	err = DB.Model(&model.Product{}).Where("name LIKE ?", "%"+name+"%").Count(&total).Error
	return total, err
}

// GetProductsByName 根据商品名模糊查询商品
func GetProductsByName(name string, page model.BasePage) (products []*model.Product, err error) {
	err = DB.Where("name LIKE ?", "%"+name+"%").Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return products, err
}

// GetProductById 根据id查询商品
func GetProductById(id uint) (product *model.Product, err error) {
	err = DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func UpdateProduct(product *model.Product) error {
	return DB.Model(&model.Product{}).Where("id = ?", product.ID).Updates(&product).Error
}
