package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", blog{})

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
