package main

import (
	"fmt"

	"letsgo/database"
	"letsgo/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool

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

	token, err := utils.GenerateToken("exampleUser")
	if err != nil {
		fmt.Println("Error generating token:", err)
	} else {
		fmt.Println("Generated Token:", token)
	}

	err = utils.VerifyToken(token)
	if err != nil {
		fmt.Println("Error verifying token:", err)
	} else {
		fmt.Println("Token is valid")
	}

	defer database.Close()
}
