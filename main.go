package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/michaelc0n/http_server_golang/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func main() {
	//setup db
	const dbPath = "db.json"

	dbClient := database.NewClient(dbPath)
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		dbClient: dbClient,
	}

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/users", apiCfg.endpointUsersHandler)
	serveMux.HandleFunc("/users/", apiCfg.endpointUsersHandler)
	serveMux.HandleFunc("/posts", apiCfg.endpointPostsHandler)
	serveMux.HandleFunc("/posts/", apiCfg.endpointPostsHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("server started on ", addr)
	log.Fatal(srv.ListenAndServe())
}

func (ac apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ac.handlerGetUser(w, r)
	case http.MethodPost:
		ac.handlerCreateUser(w, r)
	case http.MethodPut:
		ac.handlerUpdateUser(w, r)
	case http.MethodDelete:
		ac.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (ac apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ac.handlerGetPosts(w, r)
	case http.MethodPost:
		ac.handlerCreatePost(w, r)
	case http.MethodDelete:
		ac.handlerDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

type errorBody struct {
	Error string `json:"Error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call responsdWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//w.Header().Set(key, value)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}
