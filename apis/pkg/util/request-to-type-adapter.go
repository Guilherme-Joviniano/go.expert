package util

import (
	"encoding/json"
	"io"
)

func RequestToTypeAdapter[T interface{}](request io.Reader) (*T, error) {
	_type := new(T)

	err := json.NewDecoder(request).Decode(&_type)

	if err != nil {
		return _type, err
	}

	return _type, nil
}
