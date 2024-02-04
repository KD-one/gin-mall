package service

import (
	"gin-mall/dao"
	"gin-mall/serializer"
)

type CategoryService struct {
	Id uint `json:"id" form:"id"`
}

// Show 获取所有分类列表
func (service *CategoryService) Show() serializer.Response {
	categorys, err := dao.GetCategorys()
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "获取分类失败",
		}
	}

	return serializer.BuildListResponse(serializer.BuildCategorys(categorys), uint(len(categorys)))
}
