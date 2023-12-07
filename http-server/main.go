package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Payload interface {
	interface{} | []interface{}
}
type DefaultResponse[P Payload] struct {
	Errors     []Error `json:"errors"`
	Payload    P       `json:"payload"`
	StatusCode int16   `json:"status_code"`
}

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", SearchCEPHandler)

	// anonymous func
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":8080", nil)
}

func SearchCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := SearchCEP(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response = new(DefaultResponse[ViaCEP])

	response.Payload = *result
	response.StatusCode = http.StatusOK
	response.Errors = make([]Error, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func SearchCEP(cep string) (*ViaCEP, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, errors.New("erro ao fazer requisição")
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, errors.New("erro ao ler resposta")
	}

	var data ViaCEP

	err = json.Unmarshal(res, &data)

	if err != nil {
		return nil, errors.New("erro ao fazer a conversão da resposta")
	}

	return &data, nil
}
