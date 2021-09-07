package configs

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	configPath    = "../../config"
	configEnvPath = ".env"
)

type App struct {
	Name           string `json:"name"`
	Environment    string `json:"environment"`
	Debug          bool   `json:"debug"`
	GrpcPort       string `json:"grpc_port"`
	SaverChunkSize uint   `json:"saver_chunk_size"`
	SaverChunkTime uint   `json:"saver_chunk_time"`
}

type Database struct {
	DbName         string `json:"db_name"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	PoolMaxConnect int    `json:"pool_max_connect"`
}

type Kafka struct {
	Topic   string         `json:"topic"`
	Brokers []*BrokerKafka `json:"brokers"`
}

type Metrics struct {
	Prometheus Prometheus `json:"prometheus"`
	Jaeger     Jaeger     `json:"jaeger"`
}

type Prometheus struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Path string `json:"path"`
}

type Jaeger struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type BrokerKafka struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var AppConfig *App
var DatabaseConfig *Database
var KafkaConfig *Kafka
var MetricsConfig *Metrics

type funcType func(string, interface{}) (interface{}, error)

func LoadConfigs() {
	AppConfig = &App{}
	DatabaseConfig = &Database{}
	KafkaConfig = &Kafka{}
	MetricsConfig = &Metrics{}
	mapper := map[string]interface{}{
		"app":      AppConfig,
		"database": DatabaseConfig,
		"kafka":    KafkaConfig,
		"metrics":  MetricsConfig,
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
	case *Kafka:
		cnf.Topic = viper.GetString(cnf.Topic)
		for _, broker := range cnf.Brokers {
			broker.Host = viper.GetString(broker.Host)
			broker.Port = viper.GetString(broker.Port)
		}
	case *App:
		cnf.Name = viper.GetString(cnf.Name)
		cnf.Environment = viper.GetString(cnf.Environment)
		cnf.Debug = viper.GetBool("APP_DEBUG")
		cnf.GrpcPort = viper.GetString(cnf.GrpcPort)
		cnf.SaverChunkSize = viper.GetUint("SAVER_CHUNK_SIZE")
		cnf.SaverChunkTime = viper.GetUint("SAVER_CHUNK_SECOND")
	case *Metrics:
		cnf.Prometheus.Host = viper.GetString(cnf.Prometheus.Host)
		cnf.Prometheus.Port = viper.GetString(cnf.Prometheus.Port)
		cnf.Prometheus.Path = viper.GetString(cnf.Prometheus.Path)
		cnf.Jaeger.Host = viper.GetString(cnf.Jaeger.Host)
		cnf.Jaeger.Port = viper.GetString(cnf.Jaeger.Port)
	}
}
