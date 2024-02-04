package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func ShowCarousels(c *gin.Context) {
	var carousel service.CarouselService
	if err := c.ShouldBind(&carousel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := carousel.Show()
	c.JSON(200, res)
}
