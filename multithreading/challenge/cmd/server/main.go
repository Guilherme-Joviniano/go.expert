package main

import (
	"net/http"

	"github.com/Guilherme-Joviniano/multithreading-challenge/internal/handlers"
)

func main() {
	http.HandleFunc("/zipCodes", handlers.GetZipCodeHandler)
	http.ListenAndServe(":8000", nil)
}
