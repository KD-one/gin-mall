package service

import (
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	address := model.Address{
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
		UserId:  uid,
	}

	if err := dao.CreateAddress(&address); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "创建地址失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建地址成功",
	}

}

func (service *AddressService) Show(id string) serializer.Response {
	aid, _ := strconv.Atoi(id)

	address, err := dao.GetAddressById(uint(aid))
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取地址失败",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "获取地址成功",
		Data: serializer.BuildAddress(address),
	}
}

func (service *AddressService) ShowAll(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}
	addresses, err := dao.GetAddressesByUid(uid)
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "获取地址失败",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "获取地址成功",
		Data: serializer.BuildAddresses(addresses),
	}
}

func (service *AddressService) Update(c *gin.Context, id string) serializer.Response {
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

	address := model.Address{
		UserId:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}

	if err := dao.UpdateAddressById(&address, uint(aid), uid); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "更新地址失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "更新地址成功",
	}

}

func (service *AddressService) Delete(c *gin.Context, id string) serializer.Response {
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

	if err := dao.DeleteAddressById(uint(aid), uid); err != nil {
		return serializer.Response{
			Code: 50001,
			Msg:  "删除地址失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "删除地址成功",
	}
}
