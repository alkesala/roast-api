package main
import (
    "context"
    "log"
    "os"
    "github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, err
    }

    // Test connection
    var version string
    if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
        conn.Close(context.Background())
        return nil, err
    }

    log.Println(" Connected to:", version)
    return conn, nil
}
