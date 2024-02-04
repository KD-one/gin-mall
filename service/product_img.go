package service

import (
	"gin-mall/dao"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"strconv"
)

type ProductImgService struct {
}

func (p *ProductImgService) Show(id string) serializer.Response {
	pid, _ := strconv.Atoi(id)
	productImgs, err := dao.GetProductImgsById(uint(pid))
	if err != nil {
		return serializer.Response{
			Code: e.NotFound,
			Msg:  e.GetMsg(e.NotFound),
			Data: err,
		}
	}
	// 因为GetProductsImgById函数内没有进行分页操作，所以可以直接使用len(productImgs)获取总数
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
