package config

func RunConfig() error {
	return InitConfig("web_server_config","db_server_config")
}