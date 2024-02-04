package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	var address service.AddressService
	if err := c.ShouldBind(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := address.Create(c)
	c.JSON(200, res)
}

func ShowAddress(c *gin.Context) {
	var address service.AddressService
	if err := c.ShouldBind(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	aid := c.Param("id")
	// 调用service层
	res := address.Show(aid)
	c.JSON(200, res)
}

func ShowAddresses(c *gin.Context) {
	var address service.AddressService
	if err := c.ShouldBind(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := address.ShowAll(c)
	c.JSON(200, res)
}

func UpdateAddress(c *gin.Context) {
	var address service.AddressService
	if err := c.ShouldBind(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	aid := c.Param("id")
	// 调用service层
	res := address.Update(c, aid)
	c.JSON(200, res)
}

func DeleteAddress(c *gin.Context) {
	var address service.AddressService
	if err := c.ShouldBind(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	aid := c.Param("id")
	// 调用service层
	res := address.Delete(c, aid)
	c.JSON(200, res)
}
