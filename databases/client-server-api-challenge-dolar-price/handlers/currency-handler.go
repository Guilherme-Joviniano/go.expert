package handlers

import (
	"context"
	"encoding/json"
	"github.com/Guilherme-Joviniano/go-currency-api/data"
	"net/http"
)

func HandlerGetCurrencyCurrentValue(service *data.CurrencyAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter,
		r *http.Request,
	) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result, err := service.Get(context.Background(), "USD-BRL")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		currency := data.Currency{
			Value:          result.Payload.Usdbrl.Bid,
			CurrencyCodeId: 1,
		}

		err = service.Save(context.Background(), &currency)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseBody, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		_, err = w.Write(responseBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
