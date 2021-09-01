package configs

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

const (
	configPath    = "../../config"
	configEnvPath = ".env"
)

type App struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Debug       bool   `json:"debug"`
}

type Database struct {
	DbName         string `json:"db_name"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PoolMaxConnect int    `json:"pool_max_connect"`
}

var AppConfig *App
var DatabaseConfig *Database

type funcType func(string, interface{}) (interface{}, error)

func LoadConfigs() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	AppConfig = &App{}
	DatabaseConfig = &Database{}
	mapper := map[string]interface{}{
		"app":      AppConfig,
		"database": DatabaseConfig,
	}

	viper.SetConfigFile(configEnvPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Err(err)
	}

	configDir := readConfigDirectories(configPath)
	readFunc := openReadJsonFile()
	for name, config := range mapper {
		configFile := path.Join(configDir, name+".json")
		_, err := readFunc(configFile, config)
		if err != nil {
			log.Fatal().Err(err)
		}
		configParseEnv(config)
	}
}

func readConfigDirectories(configDir string) string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return path.Join(basePath, configDir)
}

func openReadJsonFile() funcType {
	return func(path string, config interface{}) (interface{}, error) {
		jsonFile, err := os.Open(path)
		if err != nil {
			return config, fmt.Errorf("not found file in path: %s", path)
		}
		if jsonFile == nil {
			return config, fmt.Errorf("not found file in path: %s", path)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		data := []byte(byteValue)
		err = json.Unmarshal(data, &config)
		if err != nil {
			return config, fmt.Errorf("Not parsed file in path: %s", path)
		}

		return config, nil
	}
}

func configParseEnv(config interface{}) {
	switch cnf := config.(type) {
	case *Database:
		cnf.DbName = viper.GetString(cnf.DbName)
		cnf.Host = viper.GetString(cnf.Host)
		cnf.Port = viper.GetString(cnf.Port)
		cnf.Username = viper.GetString(cnf.Username)
		cnf.Password = viper.GetString(cnf.Password)
		cnf.PoolMaxConnect = viper.GetInt("DB_POOL_MAX_CONNECT")
	case App:
		cnf.Name = viper.GetString(cnf.Name)
		cnf.Environment = viper.GetString(cnf.Environment)
		cnf.Debug = viper.GetBool("APP_DEBUG")
	}
}
