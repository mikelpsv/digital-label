package config

import (
	"github.com/joho/godotenv"
	app "github.com/mlplabs/app-utils"
	"os"
)

type Service struct {
	EnvProduction  bool
	EnvLocal       bool
	AppAddr        string
	AppPort        string
	DbHost         string
	DbPort         string
	DbName         string
	DbUser         string
	DbPassword     string
	KafkaHost0     string
	KafkaDataTopic string
	KafkaDataGroup string
}

func ReadEnv() *Service {
	var Cfg Service
	err := godotenv.Load()
	if err != nil {
		app.Log.Error.Println("error reading .env file")
	}

	Cfg.EnvProduction = os.Getenv("ENV") == "production"
	Cfg.EnvLocal = os.Getenv("ENV") == "local"

	Cfg.AppAddr = os.Getenv("APP_ADDR")
	Cfg.AppPort = os.Getenv("APP_PORT")

	Cfg.DbHost = os.Getenv("DB_HOST")
	Cfg.DbPort = os.Getenv("DB_PORT")
	Cfg.DbName = os.Getenv("DB_NAME")
	Cfg.DbUser = os.Getenv("DB_USER")
	Cfg.DbPassword = os.Getenv("DB_PASS")

	Cfg.KafkaHost0 = os.Getenv("KAFKA_BROKER0_HOST")
	Cfg.KafkaDataGroup = os.Getenv("KAFKA_DATA_GROUP")
	Cfg.KafkaDataTopic = os.Getenv("KAFKA_DATA_TOPIC")
	return &Cfg
}
