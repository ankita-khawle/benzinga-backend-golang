package main

import (
	"benzinga-backend-golang/controllers"
	"benzinga-backend-golang/models"
	"benzinga-backend-golang/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig() // load configs

	utils.InitLogger() // Initialize logger

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
