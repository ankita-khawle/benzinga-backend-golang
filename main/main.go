package main

import (
	"benzinga-backend-golang/controllers"
	"benzinga-backend-golang/models"
	"benzinga-backend-golang/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from environment variables
	utils.InitConfig() // Initialize logger

	utils.InitLogger()

	// Start batch processing ticker
	go utils.StartBatchProcessor()

	// Setup router and endpoints
	router := gin.Default()
	router.GET("/healthz", controllers.HealthCheck)
	router.POST("/log", controllers.HandleLog)

	// Start server
	models.Logger.Info("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		models.Logger.Fatal("Server failed to start: ", err)
	}
}
