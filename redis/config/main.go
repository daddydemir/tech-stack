package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get(key string) string {
	err := godotenv.Load("../config/local.env")
	if err != nil {
		log.Fatalln("local.env resource not found.")
	}
	return os.Getenv(key)
}
