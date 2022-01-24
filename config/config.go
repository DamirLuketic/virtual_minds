package config

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:          getEnv("VM_PORT", "8080"),
		MySQLHost:     getEnv("MYSQL_HOST", "db"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLDatabase: getEnv("MYSQL_DATABASE", "vm"),
		MySQLUser:     getEnv("MYSQL_USER", "vm"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "vm"),
		APIUser:       getEnv("API_USER", "vm"),
		APIPassword:   getEnv("API_PASSWORD", "vm"),
	}
}
