package config

import "github.com/spf13/viper"

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
	viper.SetDefault("PORT", "api_gateway:3000")
	viper.SetDefault("DB_URL", "postgres://user:secret@postgres:5432/url_redirector")
	viper.SetDefault("JWT_SECRET_KEY", "not-secret-key")
	viper.SetDefault("ISSUER", "URL-svc")
	viper.SetDefault("EXPIRATION_HOURS", 1)
	viper.SetDefault("URL_SERVICE", "url_service:50051")
	viper.SetDefault("AUTH_SERVICE", "auth_service:50052")

	viper.AutomaticEnv()
	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}

// func LoadConfig() (config Config, err error) {
// 	config.Port = os.Getenv("PORT")
// 	config.AuthSvcPort = os.Getenv("AUTH_SERVICE")
// 	config.UrlSvcPort = os.Getenv("URL_SERVICE")

// 	config.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
// 	if config.JWTSecretKey == "" {
// 		config.JWTSecretKey = "not-secret-key"
// 	}
// 	config.Issuer = os.Getenv("ISSUER")
// 	config.Port = os.Getenv("PORT")
// 	if config.Port == "" {
// 		config.Port = ":3000"
// 	}

// 	if config.AuthSvcPort == "" {
// 		config.AuthSvcPort = ":50052"
// 	}
// 	if config.UrlSvcPort == "" {
// 		config.UrlSvcPort = ":50051"
// 	}

// 	return
// }
