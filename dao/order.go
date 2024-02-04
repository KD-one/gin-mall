package dao

import "gin-mall/model"

func CreateOrder(order *model.Order) error {
	return DB.Model(&model.Order{}).Create(&order).Error
}

func GetOrderByIdAndUid(id, uid uint) (order *model.Order, err error) {
	err = DB.Where("id = ? AND user_id = ?", id, uid).First(&order).Error
	return order, err
}

func GetOrdersByUid(uid uint, page model.BasePage) (orders []*model.Order, err error) {
	err = DB.Model(&model.Order{}).Where("user_id = ?", uid).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&orders).Error
	return orders, err
}

func DeleteOrderById(oid, uid uint) error {
	return DB.Where("id = ? AND user_id = ?", oid, uid).Delete(&model.Order{}).Error
}

func GetOrderCountsByUid(uid uint) (total int64, err error) {
	err = DB.Model(&model.Order{}).Where("user_id = ?", uid).Count(&total).Error
	return
}

func UpdateOrder(order *model.Order) error {
	return DB.Model(&model.Order{}).Where("id = ?", order.ID).Updates(&order).Error
}
