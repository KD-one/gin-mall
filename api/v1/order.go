package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order service.OrderService
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := order.Create(c)
	c.JSON(200, res)
}

func ShowOrder(c *gin.Context) {
	var order service.OrderService
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	oid := c.Param("id")
	// 调用service层
	res := order.Show(c, oid)
	c.JSON(200, res)
}

func ShowOrders(c *gin.Context) {
	var order service.OrderService
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := order.ShowAll(c)
	c.JSON(200, res)
}

func DeleteOrder(c *gin.Context) {
	var order service.OrderService
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	oid := c.Param("id")
	// 调用service层
	res := order.Delete(c, oid)
	c.JSON(200, res)
}
