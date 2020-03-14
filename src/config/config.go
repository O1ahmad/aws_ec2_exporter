package config

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func CheckConfig() {
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		log.Error("*Required* key ID not set for AWS API Access.")
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		log.Error("*Required* API access secret key not.")
	}
}
