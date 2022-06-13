package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", fooApiHandler)
	mux.HandleFunc("/foo/caller", callerApiHandler)
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

func callerApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		service := req.URL.Query().Get("service")
		url := "http://" + service + ".dev.svc.cluster.local" + ":5000" + "/" + service
		responseBody := httpGetRequest(url)

		w.WriteHeader(200)
		w.Write([]byte(responseBody))
	}
}

func httpGetRequest(url string) string {
	// GET 호출
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(data)
}
