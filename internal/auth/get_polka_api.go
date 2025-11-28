package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := strings.ReplaceAll(strings.ReplaceAll(headers.Get("Authorization"), " ", ""), "ApiKey", "")
	if apiKey == "" {
		return "", errors.New("no auth header included in request")
	}
	return apiKey, nil
}
