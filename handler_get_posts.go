package main

import (
	"errors"
	"net/http"
)

func (ac apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("userEmail")
	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided"))
		return
	}
	posts, err := ac.dbClient.GetPosts(userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
