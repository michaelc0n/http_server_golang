package main

import (
	"log"
	"net/http"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}

func main() {
	http.HandleFunc("/", testHandler)

	const addr = "localhost:8080"

	serveMux := http.NewServeMux()
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(http.ListenAndServe(srv.Addr, nil))

}
