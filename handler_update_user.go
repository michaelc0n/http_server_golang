package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (ac apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Another  string `json:"another"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := ac.dbClient.UpdateUser(userEmail, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, newUser)
}
