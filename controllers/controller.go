package controllers

import (
	"net/http"
	"github.com/ankita-khawle/benzinga-backend-golang/models/model"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func handleLog(c *gin.Context) {
	var payload model.LogPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		model.logger.Warn("Invalid JSON payload: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	model.cacheMutex.Lock()
	model.cache = append(model.cache, payload)
	model.cacheMutex.Unlock()

	model.logger.Info("Payload added to cache, current size:", len(model.cache))

	// Check if batch size is reached
	if len(model.cache) >= model.batchSize {
		batch.sendBatch()
	}
}
