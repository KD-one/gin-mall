package serializer

import (
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
)

type CartSerializer struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	Num       uint   `json:"num"`
	Total     uint   `json:"total"`
	ImgPath   string `json:"img_path"`
	Price     uint   `json:"price"`
	BoosId    uint   `json:"boos_id"`
	BoosName  string `json:"boos_name"`
}

func BuildCart(cart *model.Cart) *CartSerializer {

	// 获取商家和商品信息
	boos := dao.GetUserById(cart.BoosId)
	product, _ := dao.GetProductById(cart.ProductId)

	return &CartSerializer{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Num:       cart.Num,
		Total:     cart.Total,
		ImgPath:   conf.Host + conf.Port + conf.Product + product.ImgPath,
		Price:     cart.Price,
		BoosId:    cart.BoosId,
		BoosName:  boos.Name,
	}
}

func BuildCarts(items []*model.Cart) (carts []*CartSerializer) {
	for _, item := range items {
		carts = append(carts, BuildCart(item))
	}
	return carts
}
