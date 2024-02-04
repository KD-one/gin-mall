package dao

import "gin-mall/model"

func CreateFavorite(favorite *model.Favorite) error {
	return DB.Model(&model.Favorite{}).Create(&favorite).Error
}

// GetFavoritesByUid 获取当前用户的收藏夹
func GetFavoritesByUid(uid uint) (favorites []*model.Favorite, err error) {
	err = DB.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	return favorites, err
}

func DeleteFavorite(id uint) error {
	err := DB.Model(&model.Favorite{}).Where("id = ?", id).Delete(&model.Favorite{}).Error
	return err
}

// ExistFavorite 验证商品是否已经在收藏夹中  返回true:存在,false:不存在
func ExistFavorite(favorite *model.Favorite) bool {
	var count int64
	DB.Model(&model.Favorite{}).Where("product_id = ?", favorite.ProductId).Count(&count)
	return count > 0
}
