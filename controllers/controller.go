package controllers

import (
	"benzinga-backend-golang/models"
	"benzinga-backend-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK") // returns a string
}

func HandleLog(c *gin.Context) {
	var payload models.LogPayload
	if err := c.ShouldBindJSON(&payload); err != nil { // validates against the struct, we can add a separate layer of valitions to handle proper messages here
		models.Logger.Warn("Invalid JSON payload: ", err) // validation fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	models.CacheMutex.Lock()
	models.Cache = append(models.Cache, payload) // appends new data to existing cache
	models.CacheMutex.Unlock()

	models.Logger.Info("Payload added to cache, current size:", len(models.Cache))

	// Check if batch size is reached 
	if len(models.Cache) >= models.BatchSize {
		utils.SendBatch()
	}
	c.String(http.StatusOK, "Log Captured") // returns a string

}
