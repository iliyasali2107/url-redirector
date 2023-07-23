package service_test

import (
	"log"
	"os"
	"testing"

	"name-counter-auth/pkg/config"
	"name-counter-auth/pkg/db"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	config, err := loadTestConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db.Init(config.DBUrl)
	os.Exit(m.Run())
}

func loadTestConfig() (config config.Config, err error) {
	viper.AddConfigPath("../config/envs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	viper.Unmarshal(&config)

	return config, nil
}
