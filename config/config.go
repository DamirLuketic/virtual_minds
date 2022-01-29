package config

func NewServerConfig() *Config {
	return &Config{
		ServerConfig: &ServerConfig{
			Port:        getEnv("VM_PORT", ":8080"),
			APIUser:     getEnv("API_USER", "vm"),
			APIPassword: getEnv("API_PASSWORD", "vm"),
		},
		DBConfig: &DBConfig{
			MySQLHost:     getEnv("MYSQL_HOST", "db"),
			MySQLPort:     getEnv("MYSQL_PORT", "3306"),
			MySQLDatabase: getEnv("MYSQL_DATABASE", "vm"),
			MySQLUser:     getEnv("MYSQL_USER", "vm"),
			MySQLPassword: getEnv("MYSQL_PASSWORD", "vm"),
		},
	}
}
