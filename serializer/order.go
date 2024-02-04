package serializer

import (
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
)

type OrderSerializer struct {
	Id          uint   `json:"id"`
	OrderNum    uint   `json:"order_num"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	UserId      uint   `json:"user_id"`
	AddressId   uint   `json:"address_id"`
	Num         int    `json:"num"`
	ProductId   uint   `json:"product_id"`
	BoosId      uint   `json:"boos_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Address     string `json:"address"`
	Type        uint   `json:"type"`
	ImgPath     string `json:"img_path"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) *OrderSerializer {
	return &OrderSerializer{
		Id:          order.ID,
		OrderNum:    order.OrderNum,
		CreatedAt:   order.CreatedAt.Unix(),
		UpdatedAt:   order.UpdatedAt.Unix(),
		UserId:      order.UserId,
		AddressId:   order.AddressId,
		Num:         order.Num,
		ProductId:   order.ProductId,
		BoosId:      order.BoosId,
		ProductName: product.Name,
		Price:       order.Price,
		Address:     address.Address,
		Type:        order.Type,
		ImgPath:     conf.Host + conf.Port + conf.Product + product.ImgPath,
	}
}

func BuildOrders(items []*model.Order) (orders []*OrderSerializer) {
	for _, item := range items {
		// 检验商品是否存在
		product, err := dao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}

		// 检验地址是否存在
		address, err := dao.GetAddressById(item.AddressId)
		if err != nil {
			continue
		}
		orders = append(orders, BuildOrder(item, product, address))
	}
	return orders
}
