package memwatch

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func CheckMemoryUsage() {
	var memStats runtime.MemStats

	// Retrieve memory limit from environment variable or default to 5 MB
	limitStr := os.Getenv("APP_MEMORY_LIMIT")
	var maxMemory uint64
	if limitStr != "" {
		limit, err := strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			log.Printf("Invalid APP_MEMORY_LIMIT value, defaulting to 5MB: %v\n", err)
			maxMemory = 5 * 1024 * 1024
		} else {
			maxMemory = limit * 1024 * 1024 // Convert MB to bytes
		}
	} else {
		maxMemory = 5 * 1024 * 1024 // Default to 5 MB
	}

	for {
		runtime.ReadMemStats(&memStats)
		if memStats.Alloc > maxMemory {
			log.Printf("Memory usage exceeded %v bytes, shutting down\n", maxMemory)
			os.Exit(1)
		}
		time.Sleep(10 * time.Second)
	}
}
