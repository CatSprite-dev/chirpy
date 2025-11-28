package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CatSprite-dev/chirpy/internal/auth"
	"github.com/CatSprite-dev/chirpy/internal/database"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token couldn't be access", err)
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), database.RevokeRefreshTokenParams{
		RevokedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: time.Now(),
		Token:     token,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error(), err)
	}
	respondWithJSON(w, 204, nil)
}
