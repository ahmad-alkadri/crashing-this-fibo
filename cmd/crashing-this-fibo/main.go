package main

import (
	"github.com/ahmad-alkadri/crashing-this-fibo/internal/api"
	memwatch "github.com/ahmad-alkadri/crashing-this-fibo/internal/watcher"

	"log"
	"net/http"
)

func main() {
	go memwatch.CheckMemoryUsage() // Start memory monitoring

	mux := api.NewRouter() // Setup routes using the built-in HTTP routing
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
