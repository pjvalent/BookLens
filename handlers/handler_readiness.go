package handlers

import "net/http"

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, struct{ Message string }{"server is online."})
}
