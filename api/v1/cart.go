package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	var cart service.CartService
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := cart.Create(c)
	c.JSON(200, res)
}

func ShowCart(c *gin.Context) {
	var cart service.CartService
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	cid := c.Param("id")
	// 调用service层
	res := cart.Show(cid)
	c.JSON(200, res)
}

func ShowCarts(c *gin.Context) {
	var cart service.CartService
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := cart.ShowAll(c)
	c.JSON(200, res)
}

func UpdateCart(c *gin.Context) {
	var cart service.CartService
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	cid := c.Param("id")
	// 调用service层
	res := cart.Update(c, cid)
	c.JSON(200, res)
}

func DeleteCart(c *gin.Context) {
	var cart service.CartService
	if err := c.ShouldBind(&cart); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	cid := c.Param("id")
	// 调用service层
	res := cart.Delete(c, cid)
	c.JSON(200, res)
}
