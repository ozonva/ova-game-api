package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigsLoad(t *testing.T) {
	appConfig := App{
		Name:        "APP_NAME",
		Environment: "APP_ENV",
		Debug:       false,
	}
	databaseConfig := Database{
		DbName:         "",
		Host:           "",
		Port:           "",
		Username:       "",
		Password:       "",
		PoolMaxConnect: 0,
	}

	LoadConfigs()
	assert.Equal(t, appConfig, *AppConfig)
	assert.Equal(t, databaseConfig, *DatabaseConfig)
}
