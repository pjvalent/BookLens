package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX level error: ", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{Error: message})

}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed ot marshal JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
