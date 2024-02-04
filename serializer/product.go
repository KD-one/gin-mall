package serializer

import (
	"gin-mall/conf"
	"gin-mall/model"
)

type ProductSerializer struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	ImgPath       string `json:"img_path"`
	Category      string `json:"category"`
	View          uint64 `json:"view"` // 商品浏览量
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BoosId        uint   `json:"boos_id"`
	BoosName      string `json:"boos_name"`
	BoosAvatar    string `json:"boos_avatar"`
}

func BuildProduct(product *model.Product) *ProductSerializer {
	return &ProductSerializer{
		Id:            product.ID,
		Name:          product.Name,
		Title:         product.Title,
		Info:          product.Info,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		ImgPath:       conf.Host + conf.Port + conf.Product + product.ImgPath,
		Category:      product.Category,
		View:          product.View(),
		OnSale:        product.OnSale,
		Num:           product.Num,
		BoosId:        product.BoosId,
		BoosName:      product.BoosName,
		BoosAvatar:    conf.Host + conf.Port + conf.Avatar + product.BoosAvatar,
	}
}

func BuildProducts(items []*model.Product) (products []*ProductSerializer) {
	for _, item := range items {
		products = append(products, BuildProduct(item))
	}
	return products
}
