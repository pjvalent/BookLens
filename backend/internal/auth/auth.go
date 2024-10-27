package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an api key from the headers of an http request
// Example: ---- Authorization: ApiKey {--api key here--}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header, expected format: ApiKey {--api_key--}")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header, expected format: ApiKey {--api_key--}")
	}

	return vals[1], nil

}

func GetToken(headers http.Header) (string, error) {

	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(authHeader, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header, expected format: Bearer {--bearer token--}")
	}

	if vals[0] != "Bearer" {	
		return "", errors.New("malformed first part of auth header, expected format: Bearer {--bearer token--}")
	}

	return vals[1], nil
}
