package handlers

import (
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "something went wrong")
}
