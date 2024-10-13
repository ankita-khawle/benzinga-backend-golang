package utils

import (
	"benzinga-backend-golang/models"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func SendBatch() {
	models.CacheMutex.Lock()
	if len(models.Cache) == 0 {
		models.CacheMutex.Unlock()
		return
	}

	data := models.Cache
	models.Cache = []models.LogPayload{}
	models.CacheMutex.Unlock()

	for i := 0; i < 3; i++ {
		if postData(data) {
			return
		}
		models.Logger.Warn("Failed to send batch, retrying...")
		time.Sleep(2 * time.Second)
	}

	models.Logger.Error("Failed to send batch after 3 retries, exiting.")
	os.Exit(1)
}

func postData(data []models.LogPayload) bool {
	start := time.Now()
	body, err := json.Marshal(data)
	if err != nil {
		models.Logger.Error("Failed to serialize batch: ", err)
		return false
	}
	resp, err := http.Post(models.PostEndpoint, "application/json", bytes.NewBuffer(body))
	duration := time.Since(start)

	if err != nil {
		models.Logger.Error("Failed to post batch: ", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		models.Logger.Info("Batch sent successfully. Batch size:", len(data), "Duration:", duration)
		return true
	}

	models.Logger.Error("Failed to send batch, status code:", resp.StatusCode)
	return false
}
