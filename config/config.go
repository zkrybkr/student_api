package config

import (
	"log"
	"os"
	"path/filepath"
	m "studentApi/models"

	"github.com/joho/godotenv"
)

var web_server_env *m.WebServerConfigEnv = &m.WebServerConfigEnv{}
var db_server_env *m.DBServerConfigEnv = &m.DBServerConfigEnv{}

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

}

func InitDbServerConfig() {

}
