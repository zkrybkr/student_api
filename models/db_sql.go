package models

type SQLYamlData struct {
	AddUUIDExtensionSQL string `yaml:"add_uuid_extension_sql"`
	TableCreateRolesSQL string `yaml:"table_create_roles_sql"`
	TableCreateUsersSQL string `yaml:"table_create_users_sql"`
	TableSelectSQL      string `yaml:"table_select_sql"`
	TableUsersName      string `yaml:"table_users_name"`
	TableRolesName      string `yaml:"table_roles_name"`
}
