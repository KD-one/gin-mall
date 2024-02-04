package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func PayOrder(c *gin.Context) {
	var pay service.PayService
	if err := c.ShouldBind(&pay); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := pay.Pay(c)
	c.JSON(200, res)
}
