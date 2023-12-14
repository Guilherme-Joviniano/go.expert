package pkg_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Guilherme-Joviniano/go-currency-api/types"
	"io"
	"net/http"
)

func FetchHttpThirdPartyCurrencyAPI(ctx context.Context) (*types.HttpResponse[*types.DolarAPI], error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var parsedResult = new(types.DolarAPI)

	err = json.Unmarshal(body, parsedResult)

	payload := types.HttpResponse[*types.DolarAPI]{
		Payload:    parsedResult,
		Errors:     nil,
		StatusCode: resp.StatusCode,
	}

	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout reached")
	default:
		return &payload, nil
	}
}
