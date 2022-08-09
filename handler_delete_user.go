package main

import (
	"errors"
	"net/http"
	"strings"
)

func (ac apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided to handlerDeleteUser"))
		return
	}
	err := ac.dbClient.DeleteUser(userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
