package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool  // Global variable to hold the connection pool
var ctx = context.Background() // Context for database operations

/*
   Postgres config and connection pool 
   @return: *pgxpool.Pool, error
*/
func Connect() (*pgxpool.Pool, error) {
	var err error

	pool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		fmt.Printf("Unable to ping database: %v\n", err)
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return pool, nil
}

/*
   Close the connection pool
*/
func Close() {
	if pool != nil {
		fmt.Println("Connection pool closed")
		pool.Close()
	}
}