package config

import (
	"encoding/json"
	"os"
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
	//Todo: Find better way to do this
	file, _ := os.Open("/Users/divyansh.methi/GolandProjects/bookStore/config/service-conf.json")
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
