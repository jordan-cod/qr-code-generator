package config

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func LoadDatabaseConfig() DatabaseConfig {
	host := GetEnv("DB_HOST", "localhost")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "yourpassword")
	dbName := GetEnv("DB_NAME", "yourdb")
	port := GetEnv("DB_PORT", "5432")
	sslMode := GetEnv("DB_SSLMODE", "disable")

	return DatabaseConfig{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
		SSLMode:  sslMode,
	}
}
