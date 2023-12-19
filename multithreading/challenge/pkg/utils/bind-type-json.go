package utils

import (
	"encoding/json"
	"io"
)


func BindTypeJSON[T interface{}](reader io.Reader) (*T, error) {
	bind := new(T)

	err := json.NewDecoder(reader).Decode(&bind)

	if err != nil {
		return nil, err
	}

	return bind, nil
}
