package api

import (
	"encoding/json"
	fibo "github.com/ahmad-alkadri/crashing-this-fibo/pkg/fibonacci"
	"log"
	"net/http"
	"strconv"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.URL.String())
	query := r.URL.Query()
	input, ok := query["n"]
	if !ok || len(input) != 1 {
		http.Error(w, "Missing 'n' query parameter", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(input[0])
	if err != nil {
		http.Error(w, "Invalid integer", http.StatusBadRequest)
		return
	}

	result := fibo.Fibonacci(n) // Call Fibonacci function
	response := map[string]int{"n": n, "fibonacci": result}
	json.NewEncoder(w).Encode(response)
}
