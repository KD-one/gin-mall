package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func ShowCategories(c *gin.Context) {
	var category service.CategoryService
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := category.Show()
	c.JSON(200, res)
}
