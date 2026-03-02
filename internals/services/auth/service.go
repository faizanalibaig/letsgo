package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
}

type login struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
}

var users = []user{
	{ID: "1", Name: "John Doe", Email: "john@gmail.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@yahoo.com"},
}

func Register(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, users)
}


func Login(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, users)
}