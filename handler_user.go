package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cihanalici/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	respondWithJSON(w, http.StatusOK, struct{}{})
}
