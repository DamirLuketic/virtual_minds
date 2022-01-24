package config

type ServerConfig struct {
	// Port on which vm is listening
	Port          string
	MySQLHost     string
	MySQLPort     string
	MySQLDatabase string
	MySQLUser     string
	MySQLPassword string
	APIUser       string
	APIPassword   string
}
