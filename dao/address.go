package dao

import "gin-mall/model"

func CreateAddress(address *model.Address) error {
	return DB.Model(&model.Address{}).Create(&address).Error
}

func GetAddressById(aid uint) (address *model.Address, err error) {
	err = DB.Where("id = ?", aid).First(&address).Error
	return address, err
}

func GetAddressesByUid(uid uint) (addresses []*model.Address, err error) {
	err = DB.Model(&model.Address{}).Where("user_id = ?", uid).Find(&addresses).Error
	return addresses, err
}

// UpdateAddressById 根据地址ID和用户ID唯一确定一条记录并更新地址
// 当前用户id = 2，更新地址的id = 1，但是地址id = 1的这条记录的用户id不是2，此时会将其他用户的地址修改！！
func UpdateAddressById(address *model.Address, aid, uid uint) error {
	return DB.Model(&model.Address{}).Where("id = ? AND user_id = ?", aid, uid).Updates(&address).Error
}

func DeleteAddressById(aid, uid uint) error {
	return DB.Where("id = ? AND user_id = ?", aid, uid).Delete(&model.Address{}).Error
}
