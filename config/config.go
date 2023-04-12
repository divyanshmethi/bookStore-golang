package config

import (
	"encoding/json"
	"os"
)

const (
	ConfigFilePath = "/Users/divyansh.methi/GolandProjects/bookStore/config/service-conf.json"
)

var globalConfig *AppConfig

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Data struct {
	Postgres Postgres
}

type AppConfig struct {
	Data Data
}

func InitGlobalConfig() {
	file, _ := os.Open(ConfigFilePath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&globalConfig)
	if err != nil {
		panic("Unable to read config file")
	}
}

func GetGlobalConfig() *AppConfig {
	return globalConfig
}
