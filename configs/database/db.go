package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func DatabaseConnection() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := 5432
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, errSql := sql.Open("postgres", connectionString)
	if errSql != nil {
		fmt.Println("Error connecting to configs", errSql)
	} else {
		Db = db
		fmt.Println("Connected to configs")
	}
}
