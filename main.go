package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ankita-khawle/benzinga-backend-golang/utils/common"
	"github.com/ankita-khawle/benzinga-backend-golang/utils/logs"
	"github.com/ankita-khawle/benzinga-backend-golang/controllers/controller"
	"github.com/ankita-khawle/benzinga-backend-golang/utils/batch"
	"github.com/ankita-khawle/benzinga-backend-golang/models/model"
)

func main() {
	// Load configuration from environment variables
	common.initConfig()
	// Initialize logger
	logs.initLogger()

	// Start batch processing ticker
	go batch.startBatchProcessor()

	// Setup router and endpoints
	router := gin.Default()
	router.GET("/healthz", controller.healthCheck)
	router.POST("/log", controller.handleLog)

	// Start server
	logVal := model.logger
	logVal.logger.Info("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		logVal.logger.Fatal("Server failed to start: ", err)
	}

}
