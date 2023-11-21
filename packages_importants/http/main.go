package main

import (
	"net/http"
	"io"
)

func main() {
	req, err := http.Get("https://www.google.com")
	
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	req.Body.Close()
}