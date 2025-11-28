package main

import (
	"net/http"
	"time"

	"github.com/CatSprite-dev/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token couldn't be access", err)
		return
	}
	tokenData, err := cfg.db.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}
	tokenExpired := tokenData.ExpiresAt.Before(time.Now())
	if tokenExpired || tokenData.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Token expired or revoked", nil)
		return
	}
	accessToken, err := auth.MakeJWT(tokenData.UserID, cfg.secret, auth.AccessTokenExpireTime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		Token: accessToken,
	})
}
