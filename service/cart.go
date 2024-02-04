package service

import (
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BoosId    uint `json:"boos_id" form:"boos_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

func (service *CartService) Create(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	// 检查商品是否存在
	product, err := dao.GetProductById(service.ProductId)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "要添加的商品不存在！",
		}
	}

	price, _ := strconv.Atoi(product.Price)

	cart := model.Cart{
		UserId:    uid,
		ProductId: service.ProductId,
		BoosId:    product.BoosId,
		Num:       service.Num,
		Price:     uint(price),
	}

	if err = dao.CreateCart(&cart); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "创建购物车记录失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建购物车记录成功",
	}

}

func (service *CartService) Show(id string) serializer.Response {
	cid, _ := strconv.Atoi(id)

	cart, err := dao.GetCartById(uint(cid))
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取购物车失败",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "获取购物车成功",
		Data: serializer.BuildCart(cart),
	}
}

func (service *CartService) ShowAll(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}
	carts, err := dao.GetCartsByUid(uid)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取购物车失败",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "获取购物车成功",
		Data: serializer.BuildCarts(carts),
	}
}

func (service *CartService) Update(c *gin.Context, id string) serializer.Response {
	cid, _ := strconv.Atoi(id)

	var uid uint
	if userid, ok := c.Get("userId"); ok {
		uid = userid.(uint)
	} else {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取用户id失败",
		}
	}

	// TODO: 暂时购物车的修改只能修改数量
	cart := model.Cart{
		Num: service.Num,
	}

	if err := dao.UpdateCartById(&cart, uint(cid), uid); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "更新购物车记录失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "更新购物车记录成功",
	}

}

func (service *CartService) Delete(c *gin.Context, id string) serializer.Response {
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

	if err := dao.DeleteCartById(uint(aid), uid); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "删除购物车记录失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "删除购物车记录成功",
	}
}
