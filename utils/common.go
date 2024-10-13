package utils

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Login struct {
	Time string `json:"time"`
	IP   string `json:"ip"`
}

type PhoneNumbers struct {
	Home   string `json:"home"`
	Mobile string `json:"mobile"`
}
type Meta struct {
	Logins       []Login      `json:"logins"`
	PhoneNumbers PhoneNumbers `json:"phone_numbers"`
}

var (
	cache         []LogPayload
	cacheMutex    sync.Mutex
	batchSize     int
	batchInterval int
	postEndpoint  string
	logger        = logrus.New()
)

type LogPayload struct {
	UserID    int     `json:"user_id"`
	Total     float64 `json:"total"`
	Title     string  `json:"title"`
	Meta      Meta    `json:"meta"`
	Completed bool    `json:"completed"`
}

func getEnv(key, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func initConfig() {
	var err error
	batchSize, err = strconv.Atoi(getEnv("BATCH_SIZE", "5"))
	if err != nil {
		log.Fatal("Invalid BATCH_SIZE: ", err)
	}

	batchInterval, err = strconv.Atoi(getEnv("BATCH_INTERVAL", "10"))
	if err != nil {
		log.Fatal("Invalid BATCH_INTERVAL: ", err)
	}

	postEndpoint = getEnv("POST_ENDPOINT", "http://localhost:8080")

	logger.Info("Configuration - Batch Size:", batchSize, "Batch Interval:", batchInterval, "Post Endpoint:", postEndpoint)
}
