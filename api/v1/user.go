package api

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var u service.UserService
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := u.Register()
	c.JSON(200, res)
}

func UserLogin(c *gin.Context) {
	var u service.UserService
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := u.Login()
	c.JSON(200, res)
}

func UserUpdate(c *gin.Context) {
	var u service.UserService
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	// 调用service层
	res := u.Update(c)
	c.JSON(200, res)
}

func AvatarUpdate(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}

	var u service.UserService
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 调用service层
	res := u.AvatarUpdate(c, file)
	c.JSON(200, res)
}

func SendEmail(c *gin.Context) {
	var email service.SendEmailService
	if err := c.ShouldBind(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	res := email.Send(c)
	c.JSON(200, res)
}

func ValidEmail(c *gin.Context) {
	var email service.ValidEmailService
	if err := c.ShouldBind(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	token := c.Param("token")
	res := email.Valid(token)
	c.JSON(200, res)
}

func ShowMoney(c *gin.Context) {
	var m service.MoneyService
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.LogrusObj.Errorf("参数绑定失败 %v", err)
		return
	}
	res := m.Show(c)
	c.JSON(200, res)
}
