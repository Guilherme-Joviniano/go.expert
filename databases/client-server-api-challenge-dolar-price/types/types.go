package types

type DolarAPI struct {
	Usdbrl struct {
		Bid string
	} `json:"USDBRL"`
}

type HttpResponse[T any] struct {
	StatusCode int      `json:"status_code"`
	Payload    T        `json:"payload"`
	Errors     []string `json:"errors"`
}

type DTO[T any] struct {
	Error   error
	Payload T
}
