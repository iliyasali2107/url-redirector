package config

import (
	"os"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	DBUrl           string `mapstructure:"DB_URL"`
	JWTSecretKey    string `mapstructure:"JWT_SECRET_KEY"`
	Issuer          string `mapstructure:"ISSUER"`
	ExpirationHours int    `mapstructure:"EXPIRATION_HOURS"`
	ClientPort      string `mapstructure:"CLIENT_PORT"`
}

func LoadConfig() (config Config, err error) {
	config.DBUrl = os.Getenv("POSTGRES_DNS")
	if config.DBUrl == "" {
		config.DBUrl = "postgres://user:secret@localhost:5432/url_redirector"
	}
	config.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	if config.JWTSecretKey == "" {
		config.JWTSecretKey = "not-secret-key"
	}
	config.Issuer = os.Getenv("ISSUER")
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = ":50051"
	}

	return
}
