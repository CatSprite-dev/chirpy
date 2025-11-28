package main

import (
	"net/http"

	"github.com/CatSprite-dev/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse id", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
	}

	if chirp.UserID == userID {
		err = cfg.db.DeleteChirpById(r.Context(), chirpID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
			return
		}
	} else {
		respondWithError(w, 403, "You can't delete this chirp", nil)
		return
	}

	respondWithJSON(w, 204, nil)
}
