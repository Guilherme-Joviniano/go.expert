package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second,
	}

	jsonExample := bytes.NewBuffer([]byte(`{"foo": "bar"}`))

	resp, err := client.Post("http://google.com", "application/json", jsonExample)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
