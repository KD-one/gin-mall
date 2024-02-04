package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	conf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"Content-Type", "Content-Length", "Content-Range", "Content-Disposition"},
		MaxAge:          3600 * 8,
	}
	return cors.New(conf)
}
