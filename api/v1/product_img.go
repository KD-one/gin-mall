package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func ShowProductImg(c *gin.Context) {
	var productImg service.ProductImgService
	if err := c.ShouldBind(&productImg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	pid := c.Param("id")
	// 调用service层
	res := productImg.Show(pid)
	c.JSON(200, res)
}
