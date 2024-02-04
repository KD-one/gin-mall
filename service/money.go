package service

import (
	"gin-mall/dao"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
)

type MoneyService struct {
	Key string `json:"key" form:"key"` // TODO: 解密密码的key，相当于支付密码，！当前不用！
}

func (m *MoneyService) Show(c *gin.Context) serializer.Response {
	// 获取用户信息
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}
	user := dao.GetUserById(uid)

	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: serializer.BuildMoney(user, m.Key),
	}
}
