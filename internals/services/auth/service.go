package auth

import (
	"net/http"
	"open-lms/internals/utils"

	"github.com/gin-gonic/gin"
)

type user struct {
	Email  string   `json:"email"`
    Password string   `json:"password"`
}

type login struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
}

var users = []user{
	{Email: "john@gmail.com", Password: "password123"},
	{Email: "jane@yahoo.com", Password: "password456"},
}

func Register(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = hashedPassword

	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, users)
}


func Login(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// ok := utils.ComparePassword(user.Password, )

	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, users)
}