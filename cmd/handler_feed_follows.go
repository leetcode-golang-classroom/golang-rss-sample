package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/leetcode-golang-classroom/golang-rss-sample/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	responseWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollowsByUserID(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feed follows by id: %v", err))
		return
	}

	responseWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse follow id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollowByUserIDAndFeedID(r.Context(), database.DeleteFeedFollowByUserIDAndFeedIDParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	responseWithJSON(w, http.StatusCreated, struct{}{})
}
