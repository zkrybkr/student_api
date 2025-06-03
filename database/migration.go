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

func ConnectToDB(url string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("veritabanı bağlantı URL'si parse edilirken hata oluştu: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func OpenDBConnection(dbURL string) (*pgxpool.Pool, error) {
	pool, err := ConnectToDB(dbURL)
	if err != nil {
		return nil, fmt.Errorf("veritabanına bağlanırken hata oluştu: %v", err)
	}

	return pool, nil
}

func CreateConnectURL() string {
	dbEnv := c.GetDBServerEnv()

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.Username, dbEnv.Password, dbEnv.Host, dbEnv.Port, dbEnv.DBName)
}

func CreateDBPool() (*pgxpool.Pool, error) {
	pool, err := OpenDBConnection(CreateConnectURL())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
