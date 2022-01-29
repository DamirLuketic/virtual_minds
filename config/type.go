package config

type Config struct {
	ServerConfig *ServerConfig
	DBConfig     *DBConfig
}

type ServerConfig struct {
	// Port on which vm is listening
	Port        string
	APIUser     string
	APIPassword string
}

type DBConfig struct {
	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
	MySQLUser     string
	MySQLPassword string
}
