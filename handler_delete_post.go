package main

import (
	"errors"
	"net/http"
	"strings"
)

func (ac apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	postID := strings.TrimPrefix(r.URL.Path, "/posts/")
	if postID == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no id provided to handlerDeletePost"))
		return
	}
	err := ac.dbClient.DeletePost(postID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
