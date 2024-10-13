package utils

import (
	"github.com/sirupsen/logrus"
)


func initLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.Info("Logger initialized")
}