package service

import (
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type OrderService struct {
	BoosId    uint   `json:"boos_id" form:"boos_id"`
	ProductId uint   `json:"product_id" form:"product_id"`
	UserId    uint   `json:"user_id" form:"user_id"`
	AddressId uint   `json:"address_id" form:"address_id"`
	Num       uint   `json:"num" form:"num"`
	Price     string `json:"price" form:"price"`
	OrderNum  uint   `json:"order_num" form:"order_num"` // 订单编号
	Type      uint   `json:"type" form:"type"`           // 订单的各种状态码
	model.BasePage
}

func (service *OrderService) Create(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	// 检验商家是否存在
	boos := dao.GetUserById(service.BoosId)
	if boos == nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取商家失败",
		}
	}

	// 检验商品是否存在
	_, err := dao.GetProductById(service.ProductId)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取商品失败",
		}
	}

	// 检验地址是否存在
	_, err = dao.GetAddressById(service.AddressId)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取地址失败",
		}
	}

	// 保证订单编号唯一
	orderNum := service.ProductId + service.BoosId + uint(time.Now().Unix()-1700000000)

	order := model.Order{
		UserId:    uid,
		ProductId: service.ProductId,
		BoosId:    service.BoosId,
		AddressId: service.AddressId,
		Num:       int(service.Num),
		Price:     service.Price,
		OrderNum:  orderNum,
		Type:      service.Type,
	}

	if err = dao.CreateOrder(&order); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "创建订单失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建订单成功",
	}

}

func (service *OrderService) Show(c *gin.Context, id string) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}
	oid, _ := strconv.Atoi(id)

	// 根据订单编号和用户id唯一确定一个订单
	order, err := dao.GetOrderByIdAndUid(uint(oid), uid)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取订单失败",
		}
	}

	// 根据订单地址id获取地址信息
	address, err := dao.GetAddressById(order.AddressId)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取地址失败",
		}
	}

	// 根据订单商品id获取商品信息
	product, err := dao.GetProductById(order.ProductId)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取商品失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "获取订单成功",
		Data: serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) ShowAll(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orders, err := dao.GetOrdersByUid(uid, service.BasePage)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取订单失败",
		}
	}

	total, err := dao.GetOrderCountsByUid(uid)
	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(total))
}

func (service *OrderService) Delete(c *gin.Context, id string) serializer.Response {
	aid, _ := strconv.Atoi(id)

	var uid uint
	if userid, ok := c.Get("userId"); ok {
		uid = userid.(uint)
	} else {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取用户id失败",
		}
	}

	if err := dao.DeleteOrderById(uint(aid), uid); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "删除订单失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "删除订单成功",
	}
}
