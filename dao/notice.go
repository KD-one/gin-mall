package dao

import "gin-mall/model"

func GetNoticeById(id uint) *model.Notice {
	var notice model.Notice
	DB.Model(&model.Notice{}).Where("id = ?", id).Find(&notice)
	return &notice
}
