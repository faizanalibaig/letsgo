package user

import (
	"go-backend/configs/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int    `json:"id"`
	NAME   string `json:"name"`
	EMAIL  string `json:"email"`
	GENDER string `json:"gender"`
}

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	_, err = database.Db.Exec("INSERT INTO users(id, name, email, gender) VALUES($1, $2, $3, $4)", user.ID, user.NAME, user.EMAIL, user.GENDER)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error while inserting user ",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

func GetUsers(c *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get users",
			"error":   err.Error(),
		})
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.NAME, &user.EMAIL, &user.GENDER); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to scan user",
				"error":   err.Error(),
			})
			return
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Row iteration error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"users":   users,
	})
}
