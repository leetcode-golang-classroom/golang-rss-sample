package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/leetcode-golang-classroom/golang-rss-sample/internal/database"
)

func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}
	responseWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiconfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiconfig.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
