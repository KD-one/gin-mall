package middleware

import (
	"gin-mall/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func CheckToken(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "请求未携带token",
		})
		c.Abort()
		return
	}

	_, claims, err := util.ParseToken(tokenString[7:])

	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "token解析失败",
		})
		c.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "token已过期",
		})
		c.Abort()
		return
	}

	c.Set("userId", claims.UserId)
	c.Next()
}
