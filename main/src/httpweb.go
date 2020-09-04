package main

import (
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", webHelloWorld)
	http.ListenAndServe(":8000", nil)
}

func webHelloWorld(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hellow World"))
}
func HttpServerNew() {
	http.HandleFunc("/", webHelloWorld)
	http.ListenAndServe(":8010", nil)
}
