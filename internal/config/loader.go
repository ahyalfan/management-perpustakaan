package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load() // otomatis env file
	if err != nil {
		log.Fatal("Error loading config env file: ", err.Error())
	}
	expInt, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		log.Fatal("Error convert config env file: ", err.Error())
	}

	return &Config{
		Server: Server{
			Host:  os.Getenv("SERVER_HOST"),
			Port:  os.Getenv("SERVER_PORT"),
			Asset: os.Getenv("SERVER_ASSET_URL"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: expInt,
		},
		Storage: Storage{
			BasePath: os.Getenv("STORAGE_PATH"),
		},
	}
}
