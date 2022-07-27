package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/michaelc0n/http_server_golang/internal/database"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"mr":"si"}`))
}

func main() {
	//setup db
	c := database.NewClient("db.json")
	fmt.Println(c.Path)
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")

	user, err := c.CreateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user created", user)

	serveMux := http.NewServeMux()

	//start server
	http.HandleFunc("/", testHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(http.ListenAndServe(srv.Addr, nil))

}
