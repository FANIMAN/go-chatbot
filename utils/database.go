package utils

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDatabase() {
    databaseUrl := os.Getenv("DATABASE_URL")
    if databaseUrl == "" {
        log.Fatal("DATABASE_URL environment variable is required")
    }

    config, err := pgxpool.ParseConfig(databaseUrl)
    if err != nil {
        log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
    }

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    DB = pool
    fmt.Println("Connected to the database!")
}
