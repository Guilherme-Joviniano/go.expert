package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Guilherme-Joviniano/multithreading-challenge/internal/entity"
	infra_http "github.com/Guilherme-Joviniano/multithreading-challenge/internal/infra/http"
	"github.com/Guilherme-Joviniano/multithreading-challenge/pkg/dto"
	"github.com/go-chi/chi/v5"
)

func GetZipCodeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	zipCode := chi.URLParam(r, "zipCode")

	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viaCepResultChannel := make(chan *dto.ZipCodeDetails)
	viaBrasilAPICepResultChannel := make(chan *dto.ZipCodeDetails)

	go func() {
		result, err := infra_http.GetHTTP[entity.ViaCep](entity.ViaCepUrl + zipCode + "/json/")

		if err != nil {
			return
		}

		viaCepResultChannel <- dto.NewZipCodeDetails(result.Logradouro, result.Cep, result.Localidade, result.Bairro, result.Uf)
	}()

	go func() {
		result, err := infra_http.GetHTTP[entity.BrasilAPICep](entity.BrasilAPICepUrl + zipCode)

		if err != nil {
			return
		}

		viaBrasilAPICepResultChannel <- dto.NewZipCodeDetails(result.Cep, result.City, result.Neighborhood, result.Street, result.State)
	}()

	select {
	case value := <-viaBrasilAPICepResultChannel:

		response, err := json.Marshal(&value)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	case value := <-viaCepResultChannel:
		response, err := json.Marshal(&value)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	case <-time.After(time.Second * 1):
		fmt.Print("Timeout Reached")
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}
}
