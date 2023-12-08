package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Request initiated")
	defer log.Println("Request closed")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processed with success")
		w.Write([]byte("Request processed with success"))
		return
	case <-ctx.Done():
		log.Println("Request cancelled")
		return
	}
}
