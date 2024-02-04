package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var product service.ProductService
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := product.Create(c, files)
	c.JSON(200, res)
}

func ShowProducts(c *gin.Context) {
	var product service.ProductService
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := product.Show(c)
	c.JSON(200, res)
}

func SearchProduct(c *gin.Context) {
	var product service.ProductService
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := product.Search(c)
	c.JSON(200, res)
}

func ShowProduct(c *gin.Context) {
	var product service.ProductService
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	id := c.Param("id")
	// 调用service层
	res := product.ProductInfo(c, id)
	c.JSON(200, res)
}
