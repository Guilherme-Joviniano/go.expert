package main

import (
	"github.com/Guilherme-Joviniano/go-currency-api/routes"
	"net/http"
)

func main() {
	currencyMux := routes.CurrencyRoute()

	err := http.ListenAndServe(":8080", currencyMux)

	if err != nil {
		panic(err)
	}
}
