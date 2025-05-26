package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
    dbPool *pgxpool.Pool
    once   sync.Once
)

func InitDB() *pgxpool.Pool {
    once.Do(func() {
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }

        dbUrl := os.Getenv("DATABASE_URL")
        if dbUrl == "" {
            log.Fatal("DATABASE_URL not set in environment")
        }

        pool, err := pgxpool.New(context.Background(), dbUrl)
        if err != nil {
            log.Fatalf("Unable to create connection pool: %v", err)
        }

        dbPool = pool
        fmt.Println("✅ Database connected")
    })

    return dbPool
}