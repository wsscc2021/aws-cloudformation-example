package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fooApiHandler)
	mux.HandleFunc("/health", healthApiHandler)
	http.ListenAndServe(":5000", handlers.CombinedLoggingHandler(os.Stdout, mux))
}

func healthApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func fooApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.WriteHeader(200)
		w.Write([]byte("This is foo webapp!"))
	}
}
