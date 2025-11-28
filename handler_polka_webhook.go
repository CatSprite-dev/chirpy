package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/CatSprite-dev/chirpy/internal/auth"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func (cfg *apiConfig) handlerUserUpgradeByPolka(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load()
	polkaKey := os.Getenv("POLKA_KEY")
	if polkaKey == "" {
		log.Fatal("POLKA_KEY is not set")
	}
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "no auth header included in request", err)
		return
	}
	if apiKey != polkaKey {
		respondWithError(w, 401, "Incorrect api key", err)
		return
	}

	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, 204, nil)
		return
	}

	err = cfg.db.UpgradeUser(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}
	respondWithJSON(w, 204, nil)
}
