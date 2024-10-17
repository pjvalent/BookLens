package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	header := make(http.Header)
	header.Set("Authorization", "ApiKey testApiKey")

	_, err := GetApiKey(header)

	if err != nil {
		t.Errorf("Failed to parse apiKey")
	}

	header.Set("Authorization", "")

	_, err = GetApiKey(header)

	if err == nil {
		t.Errorf("GetApiKey should have failed but it did not")
	}
}

func TestGetToken(t *testing.T) {

	header := make(http.Header)
	header.Set("Authorization", "Bearer testBearerToken")

	_, err := GetToken(header)

	if err != nil {
		t.Errorf("Failed to parse apiKey")
	}

	header.Set("Authorization", "")

	_, err = GetToken(header)

	if err == nil {
		t.Errorf("Get token should have failed but it did not")
	}

}
