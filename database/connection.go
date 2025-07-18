package database

import (
	"context"
	"fmt"
	c "studentApi/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnData struct {
	Pool *pgxpool.Pool
}

var DBConn *DBConnData

func createConnURL() string {
	dbEnv := c.GetDBServerEnv()
	return fmt.Sprintf(dbEnv.Url, dbEnv.Driver, dbEnv.Username, dbEnv.Password, dbEnv.Host, dbEnv.Port, dbEnv.DBName)
}

func createDBPool() error {
	config, err := pgxpool.ParseConfig(createConnURL())
	if err != nil {
		return fmt.Errorf("bağlantı url parse error: %v", err)
	}

	DBConn.Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("veritabanına bağlanırken hata oluştu: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = DBConn.Pool.Ping(ctx); err != nil {
		return fmt.Errorf("db connection error: %v", err)
	}

	return nil
}
