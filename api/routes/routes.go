package routes

import (
	"bionic.pro/reports-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/reports", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Report endpoint"})
	})
}
