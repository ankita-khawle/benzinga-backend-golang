package controllers

import (
	"benzinga-backend-golang/models"
	"net/http"
	"sync"

	// "benzinga-backend-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	cache         []models.LogPayload
	cacheMutex    sync.Mutex
	batchSize     int
	batchInterval int
	postEndpoint  string
	logger        = logrus.New()
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func HandleLog(c *gin.Context) {
	var payload models.LogPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Warn("Invalid JSON payload: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	cacheMutex.Lock()
	cache = append(cache, payload)
	cacheMutex.Unlock()

	logger.Info("Payload added to cache, current size:", len(cache))

	// Check if batch size is reached
	if len(cache) >= batchSize {
		// batch.sendBatch()

	}
}
