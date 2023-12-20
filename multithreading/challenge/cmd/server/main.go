package main

import (
	"net/http"

	"github.com/Guilherme-Joviniano/multithreading-challenge/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	
	router.Get("/zipCodes/{zipCode}", handlers.GetZipCodeHandler)
	
	http.ListenAndServe(":8000", router)
}
