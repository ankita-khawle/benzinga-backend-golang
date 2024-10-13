package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func startBatchProcessor() {
	ticker := time.NewTicker(time.Duration(batchInterval) * time.Second)
	for range ticker.C {
		sendBatch()
	}
}

func sendBatch() {
	cacheMutex.Lock()
	if len(cache) == 0 {
		cacheMutex.Unlock()
		return
	}

	data := cache
	cache = []LogPayload{}
	cacheMutex.Unlock()

	for i := 0; i < 3; i++ {
		if postData(data) {
			return
		}
		logger.Warn("Failed to send batch, retrying...")
		time.Sleep(2 * time.Second)
	}

	logger.Error("Failed to send batch after 3 retries, exiting.")
	os.Exit(1)
}

func postData(data []LogPayload) bool {
	start := time.Now()
	body, err := json.Marshal(data)
	if err != nil {
		logger.Error("Failed to serialize batch: ", err)
		return false
	}

	resp, err := http.Post(postEndpoint, "application/json", bytes.NewBuffer(body))
	duration := time.Since(start)

	if err != nil {
		logger.Error("Failed to post batch: ", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		logger.Info("Batch sent successfully. Batch size:", len(data), "Duration:", duration)
		return true
	}

	logger.Error("Failed to send batch, status code:", resp.StatusCode)
	return false
}
