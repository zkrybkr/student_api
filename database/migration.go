package database

import (
	"context"
	"fmt"
	"log"
	c "studentApi/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool struct {
	Pool *pgxpool.Pool
}

var DBEngine *DBPool

func InitDBEngine(pool *pgxpool.Pool) {
	if pool == nil {
		log.Fatal("Database connection is nil! Check GetDBPool function...")
	}

	DBEngine = &DBPool{Pool: pool}
}

func CreateDBPool() (*pgxpool.Pool, error) {
	dbEnv := c.GetDBServerEnv()
	url := fmt.Sprintf(dbEnv.Url, dbEnv.Driver, dbEnv.Username, dbEnv.Password, dbEnv.Host, dbEnv.Port, dbEnv.DBName)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("veritabanı bağlantı URL'si parse edilirken hata oluştu: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("veritabanına bağlanırken hata oluştu: %v", err)
	}

	return pool, nil
}
