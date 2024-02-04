package dao

import "gin-mall/model"

func CreateCart(cart *model.Cart) error {
	return DB.Model(&model.Cart{}).Create(&cart).Error
}

func GetCartById(id uint) (cart *model.Cart, err error) {
	err = DB.Where("id = ?", id).First(&cart).Error
	return cart, err
}

func GetCartsByUid(uid uint) (carts []*model.Cart, err error) {
	err = DB.Model(&model.Cart{}).Where("user_id = ?", uid).Find(&carts).Error
	return carts, err
}

// UpdateCartById 根据购物车ID和用户ID唯一确定一条记录并更新
// 当前用户id = 2，更新购物车的id = 1，但是购物车id = 1的这条记录的用户id不是2，此时会将其他用户的购物车修改！！
func UpdateCartById(cart *model.Cart, cid, uid uint) error {
	return DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cid, uid).Updates(&cart).Error
}

func DeleteCartById(cid, uid uint) error {
	return DB.Where("id = ? AND user_id = ?", cid, uid).Delete(&model.Cart{}).Error
}
