package main

import (
	"fmt"
	"net/http"

	"github.com/cihanalici/blog-aggregator/internal/auth"
	"github.com/cihanalici/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			// errors with status code 403 are client errors
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Error getting API key: %v", err))
			return
		}

		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
