package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cihanalici/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apicfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apicfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apicfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting posts for user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
