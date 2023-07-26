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
	UrlSvcPort      string `mapstructure:"URL_SERVICE"`
	AuthSvcPort     string `mapstructure:"AUTH_SERVICE"`
}

func LoadConfig() (config Config, err error) {
	config.Port = os.Getenv("PORT")
	config.AuthSvcPort = os.Getenv("AUTH_SERVICE")
	config.UrlSvcPort = os.Getenv("URL_SERVICE")

	config.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	if config.JWTSecretKey == "" {
		config.JWTSecretKey = "not-secret-key"
	}
	config.Issuer = os.Getenv("ISSUER")
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = ":3000"
	}

	if config.AuthSvcPort == "" {
		config.AuthSvcPort = ":50052"
	}
	if config.UrlSvcPort == "" {
		config.UrlSvcPort = ":50051"
	}

	return
}
