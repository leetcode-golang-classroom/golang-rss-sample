package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/leetcode-golang-classroom/golang-rss-sample/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	responseWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}
	responseWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
