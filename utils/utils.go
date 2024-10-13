package utils

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"benzinga-backend-golang/models"
)
func getEnv(key, defaultValue string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func InitConfig(){
	var err error
	models.BatchSize, err = strconv.Atoi(getEnv("BATCH_SIZE", "5"))
	if err != nil {
		log.Fatal("Invalid BATCH_SIZE: ", err)
	}

	models.BatchInterval, err = strconv.Atoi(getEnv("BATCH_INTERVAL", "10"))
	if err != nil {
		log.Fatal("Invalid BATCH_INTERVAL: ", err)
	}

	models.PostEndpoint = getEnv("POST_ENDPOINT", "http://localhost:8080")

	models.Logger.Info("Configuration - Batch Size:", models.BatchSize, "Batch Interval:", models.BatchInterval, "Post Endpoint:", models.PostEndpoint)
}
