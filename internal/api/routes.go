package api

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/fib", FibHandler) // Register the handler for the /fib endpoint
	return mux
}
