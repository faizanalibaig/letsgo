package main

import (
	"fmt"
	"letsgo/database"

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

	defer database.Close()
}
