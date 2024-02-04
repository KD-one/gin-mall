package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func CreateFavorite(c *gin.Context) {
	var favorite service.FavoriteService
	if err := c.ShouldBind(&favorite); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := favorite.Create(c)
	c.JSON(200, res)
}

func ShowFavorites(c *gin.Context) {
	var favorite service.FavoriteService
	if err := c.ShouldBind(&favorite); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := favorite.Show(c)
	c.JSON(200, res)
}

func DeleteFavorite(c *gin.Context) {
	var favorite service.FavoriteService
	if err := c.ShouldBind(&favorite); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	id := c.Param("id")
	// 调用service层
	res := favorite.Delete(c, id)
	c.JSON(200, res)
}
