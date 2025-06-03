package config

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	m "studentApi/models"

	"github.com/joho/godotenv"
)

var web_server_env *m.WebServerConfigEnv = &m.WebServerConfigEnv{}
var db_server_env *m.DBServerConfigEnv = &m.DBServerConfigEnv{}

func GetDBServerEnv() m.DBServerConfigEnv {
	dbConfigEnv := reflect.ValueOf(db_server_env).Elem()

	for i := 0; i < dbConfigEnv.NumField(); i++ {
		fieldValue := dbConfigEnv.Field(i)
		if fieldValue.Kind() != reflect.String && fieldValue.String() == "" {
			InitDbServerConfig()
		}
	}

	return *db_server_env
}

func GetWebServerEnv() m.WebServerConfigEnv {
	webConfigEnv := reflect.ValueOf(web_server_env).Elem()

	for i := 0; i < webConfigEnv.NumField(); i++ {
		fieldValue := webConfigEnv.Field(i)

		if fieldValue.Kind() != reflect.String && fieldValue.String() == "" {
			InitWebServerConfig()
		}
	}

	return *web_server_env
}

func InitConfig(configType ...string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	envPath := filepath.Join(dir, "environment", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println(".env dosyası yüklenirken hata oluştu.")
		return err
	}

	for _, value := range configType {
		switch value {
		case "web_server_config":
			InitWebServerConfig()
		case "db_server_config":
			InitDbServerConfig()
		default:
			log.Println("belirsiz config type...")
		}
	}

	return nil
}

func InitWebServerConfig() {
	externalEnv := m.WebServerConfigEnv{
		Port:         os.Getenv("SERVER_PORT"),
		AllowMethods: strings.Split(os.Getenv("ALLOWMETHODS"), ","),
		AllowOrigins: strings.Split(os.Getenv("ALLOWORIGINS"), ","),
		AllowHeaders: strings.Split(os.Getenv("ALLOWHEADERS"), ","),
	}

	web_server_env = &externalEnv
}

func InitDbServerConfig() {
	externalEnv := m.DBServerConfigEnv{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db_server_env = &externalEnv
}
