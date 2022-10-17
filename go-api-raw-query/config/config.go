package config

import (
	"fico_ar/infrastructure/database"
	"fico_ar/infrastructure/shared/constant"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println(err)
		err = fmt.Errorf(constant.ErrLoadENV, err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Println(err)
		err = fmt.Errorf(constant.ErrConvertStringToInt, err)
		return
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.DatabaseConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	return
}
