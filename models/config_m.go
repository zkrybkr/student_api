package models

type DBServerConfigEnv struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type WebServerConfigEnv struct {
	Port         string
	AllowMethods []string
	AllowOrigins []string
	AllowHeaders []string
}
