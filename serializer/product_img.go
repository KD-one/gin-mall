package serializer

import (
	"gin-mall/conf"
	"gin-mall/model"
)

type ProductImgSerializer struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(productImg *model.ProductImg) *ProductImgSerializer {
	return &ProductImgSerializer{
		ProductId: productImg.ProductId,
		ImgPath:   conf.Host + conf.Port + conf.Product + productImg.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []*ProductImgSerializer) {
	for _, item := range items {
		productImgs = append(productImgs, BuildProductImg(item))
	}
	return productImgs
}
