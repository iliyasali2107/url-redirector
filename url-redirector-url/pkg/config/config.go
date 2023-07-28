package config

import (
	"github.com/spf13/viper"
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
	viper.SetDefault("PORT", "url_service:50051")
	viper.SetDefault("DB_URL", "postgres://user:secret@postgres:5432/url_redirector")
	viper.SetDefault("JWT_SECRET_KEY", "not-secret-key")
	viper.SetDefault("ISSUER", "URL-svc")
	viper.SetDefault("EXPIRATION_HOURS", 1)

	viper.AutomaticEnv()
	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}

// func LoadConfig() (config Config, err error) {
// 	config.DBUrl = os.Getenv("POSTGRES_DNS")
// 	if config.DBUrl == "" {
// 		config.DBUrl = "postgres://user:secret@localhost:5432/url_redirector"
// 	}
// 	config.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
// 	if config.JWTSecretKey == "" {
// 		config.JWTSecretKey = "not-secret-key"
// 	}
// 	config.Issuer = os.Getenv("ISSUER")
// 	config.Port = os.Getenv("PORT")
// 	if config.Port == "" {
// 		config.Port = ":50051"
// 	}

// 	return
// }
