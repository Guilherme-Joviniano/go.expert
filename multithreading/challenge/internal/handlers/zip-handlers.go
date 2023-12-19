package handlers

import (
	"net/http"
	"time"

	"github.com/Guilherme-Joviniano/multithreading-challenge/internal/entity"
	infra_http "github.com/Guilherme-Joviniano/multithreading-challenge/internal/infra/http"
	"github.com/Guilherme-Joviniano/multithreading-challenge/pkg/dto"
)

func GetZipCodeHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	zipCode := r.URL.Query().Get("zipCode")

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

		viaCepResultChannel <- dto.NewZipCodeDetails(result.Localidade)
	}()

	go func() {
		result, err := infra_http.GetHTTP[entity.BrasilAPICep](entity.BrasilAPICepUrl + zipCode)

		if err != nil {
			return
		}

		viaBrasilAPICepResultChannel <- dto.NewZipCodeDetails(result.Street)
	}()

	select {
	case value := <-viaBrasilAPICepResultChannel:
		println(value.Address)
		w.Write([]byte("From Brasil API"))
		return
	case value := <-viaCepResultChannel:
		println(value.Address)
		w.Write([]byte("From ViaCep API"))
		return
	case <-time.After(time.Second * 100):
		println("timeout")
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}
}
