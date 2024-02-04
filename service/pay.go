package service

import (
	"fmt"
	"gin-mall/dao"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PayService struct {
	OrderId   uint   `json:"order_id" form:"order_id"`
	OrderNum  int    `json:"order_num" form:"order_num"`
	Price     string `json:"price" form:"price"`
	Num       int    `json:"num" form:"num"`
	Key       string `json:"key" form:"key"`
	PayTime   string `json:"pay_time" form:"pay_time"`
	ProductId uint   `json:"product_id" form:"product_id"`
	BoosId    uint   `json:"boos_id" form:"boos_id"`
	BoosName  string `json:"boos_name" form:"boos_name"`
}

func (service *PayService) Pay(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	// 1、开启支付事务
	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order, err := dao.GetOrderByIdAndUid(service.OrderId, uid)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "订单不存在",
			Error: err.Error(),
		}
	}
	// 0.未支付  1.已支付
	if order.Type == 1 {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "订单已支付",
			Error: err.Error(),
		}
	}

	// 2、计算出需要支付的价格
	price, _ := strconv.ParseFloat(order.Price, 64)
	money := price * float64(order.Num)

	// 获取用户信息（用户余额）
	user := dao.GetUserById(uid)
	userMoney, _ := strconv.ParseFloat(user.Money, 64)

	// 判断用户余额是否充足
	if userMoney < money {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "余额不足",
			Error: err.Error(),
		}
	}

	// 3、将用户支付后的余额更新
	finallyMoney := fmt.Sprintf("%f", userMoney-money)
	user.Money = finallyMoney
	if err = dao.UpdateUser(user); err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "更新用户余额失败",
			Error: err.Error(),
		}
	}

	// 获取商家信息
	boos := dao.GetUserById(service.BoosId)

	// 4、更新商家收到利润后的余额
	boosMoney, _ := strconv.ParseFloat(boos.Money, 64)
	finallyBoosMoney := fmt.Sprintf("%f", boosMoney+money)
	boos.Money = finallyBoosMoney
	if err = dao.UpdateUser(boos); err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "更新商家余额失败",
			Error: err.Error(),
		}
	}

	// 5、减少相应的商品数目
	product, err := dao.GetProductById(service.ProductId)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "商品信息获取错误！",
			Error: err.Error(),
		}
	}
	product.Num -= order.Num

	if err = dao.UpdateProduct(product); err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "商品信息更新失败！",
			Error: err.Error(),
		}
	}

	// 方法一： 6、更新订单状态   可以将旧订单备份在数据库中
	order.Type = 1
	if err = dao.UpdateOrder(order); err != nil {
		tx.Rollback()
		return serializer.Response{
			Code:  500,
			Msg:   "更新订单状态失败！",
			Error: err.Error(),
		}
	}

	//// 方法二： 6、删除订单
	//if err = dao.DeleteOrderById(service.OrderId, uid); err != nil {
	//	tx.Rollback()
	//	return serializer.Response{
	//		Code:  500,
	//		Msg:   "订单删除失败！",
	//		Error: err.Error(),
	//	}
	//}

}
