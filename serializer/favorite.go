package serializer

import (
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
)

type FavoriteSerializer struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BoosId        uint   `json:"boos_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite) *FavoriteSerializer {

	// 获取商家和商品信息
	boos := dao.GetUserById(favorite.BoosId)
	product, _ := dao.GetProductById(favorite.ProductId)

	// 构建序列化对象
	return &FavoriteSerializer{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		Name:          product.Name,
		Category:      product.Category,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.Port + conf.Product + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BoosId:        boos.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(items []*model.Favorite) (favorites []*FavoriteSerializer) {
	for _, item := range items {
		favorites = append(favorites, BuildFavorite(item))
	}
	return favorites
}
