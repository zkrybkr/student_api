package models

type DBServerConfigEnv struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	Driver   string
	Url      string
}

type WebServerConfigEnv struct {
	Port         string
	AllowMethods []string
	AllowOrigins []string
	AllowHeaders []string
}
