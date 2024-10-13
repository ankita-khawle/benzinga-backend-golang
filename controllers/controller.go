package controllers

import (
	"benzinga-backend-golang/models"
	"benzinga-backend-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func HandleLog(c *gin.Context) {
	var payload models.LogPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		models.Logger.Warn("Invalid JSON payload: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	models.CacheMutex.Lock()
	models.Cache = append(models.Cache, payload)
	models.CacheMutex.Unlock()

	models.Logger.Info("Payload added to cache, current size:", len(models.Cache))

	// Check if batch size is reached
	if len(models.Cache) >= models.BatchSize {
		utils.SendBatch()
	}
}
