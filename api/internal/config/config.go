package config

import (
	"context"
	"log"

	"github.com/TheNestDevs/SnTx/api/internal/gen/appconfig"
)

type AppConfig struct {
	AppConfig *appconfig.AppConfig
}

func NewAppConfig() *AppConfig {
	config, err := appconfig.LoadFromPath(context.TODO(), "internal/config/config.pkl")
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		AppConfig: config,
	}
}
