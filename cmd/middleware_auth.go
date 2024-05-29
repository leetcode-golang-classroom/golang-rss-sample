package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/leetcode-golang-classroom/golang-rss-sample/internal/auth"
	"github.com/leetcode-golang-classroom/golang-rss-sample/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		if user.ID == uuid.Nil {
			responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user with api key: %v", apiKey))
			return
		}
		handler(w, r, user)
	}
}
