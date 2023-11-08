package config

import (
	"log"
	"os"
)

var (
	env string
)

func init() {
	env = os.Getenv("ENV")
	env = "development"
	if env == "" {
		log.Fatal("ENV environment variable is empty")
	}
}

func IsDevelopment() bool {
	return env == "development"
}
