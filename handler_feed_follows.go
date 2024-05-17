package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cihanalici/blog-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apicfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apicfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
