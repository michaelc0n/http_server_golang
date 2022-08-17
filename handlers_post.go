package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
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

func (ac apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	newPost, err := ac.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, newPost)
}

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
