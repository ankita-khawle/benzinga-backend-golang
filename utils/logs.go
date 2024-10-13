package utils

import (
	"github.com/sirupsen/logrus"
	"benzinga-backend-golang/models"
)

func InitLogger() {
	models.Logger.SetFormatter(&logrus.JSONFormatter{})
	models.Logger.SetLevel(logrus.InfoLevel)
	models.Logger.Info("Logger initialized")
}