package service

import (
	"gin-mall/dao"
	"gin-mall/serializer"
)

type CarouselService struct {
}

// Show 展示轮播图
func (service *CarouselService) Show() serializer.Response {
	carousels, err := dao.GetCarousels()
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "获取轮播图失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
