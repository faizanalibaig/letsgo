package main

import (
	"go-backend/configs/database"
	"go-backend/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.DatabaseConnection()

	r.GET("/health-check", user.Ping)
	r.POST("/add-user", user.AddUser)

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
