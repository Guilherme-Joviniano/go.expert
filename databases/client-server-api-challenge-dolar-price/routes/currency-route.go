package routes

import (
	"github.com/Guilherme-Joviniano/go-currency-api/data"
	"github.com/Guilherme-Joviniano/go-currency-api/database"
	"github.com/Guilherme-Joviniano/go-currency-api/handlers"
	"net/http"
)

func CurrencyRoute() *http.ServeMux {
	server := http.NewServeMux()

	connection, err := database.Connect()

	if err != nil {
		panic(err)
	}

	currencyService := data.NewCurrencyAdapter(connection.DB)

	server.HandleFunc("/cotacao", handlers.HandlerGetCurrencyCurrentValue(currencyService))

	return server
}
