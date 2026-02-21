package main

import (
	"fmt"
	"letsgo/database"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool
var secret = []byte("your-secret-key")

func main() {
	var err error
	err = godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	pool, err = database.Connect()
	if err != nil {
		panic(err)
	}

	token, err := GenerateToken("exampleUser")
	if err != nil {
		fmt.Println("Error generating token:", err)
	} else {
		fmt.Println("Generated Token:", token)
	}

	defer database.Close()
}

func GenerateToken (user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"time": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}