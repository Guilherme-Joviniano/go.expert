package util

import (
	"encoding/json"
	"io"
)

func RequestToTypeAdapter[T interface{}](request io.Reader) (*T, error) {
	bind := new(T)

	err := json.NewDecoder(request).Decode(&bind)

	if err != nil {
		return nil, err
	}

	return bind, nil
}
