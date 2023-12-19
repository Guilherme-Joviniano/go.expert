package infra_http

import (
	"net/http"

	"github.com/Guilherme-Joviniano/multithreading-challenge/pkg/utils"
)

func GetHTTP[T interface{}](url string) (*T, error) {
	response, err := http.DefaultClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	value, err := utils.BindTypeJSON[T](response.Body)

	if err != nil {
		return nil, err
	}

	return value, nil
}
