package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"open-lms/database"
	"open-lms/internals/services/auth"
)


func main() {
	r := gin.Default()

	database.Connect()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "backend is running",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		auth.Register(c)
	})

	r.Run(":9090")
}
