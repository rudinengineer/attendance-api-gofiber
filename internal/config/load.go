package config

import "os"

func Load() *Config {
	return &Config{
		Server: Server{
			Host:     os.Getenv("SERVER_HOST"),
			Port:     os.Getenv("SERVER_PORT"),
			Timezone: os.Getenv("TIMEZONE"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
		JWT: JWT{
			SecretKey:   os.Getenv("JWT_SECRET_KEY"),
			ExpiredTime: os.Getenv("JWT_EXPIRED_TIME"),
		},
	}
}
