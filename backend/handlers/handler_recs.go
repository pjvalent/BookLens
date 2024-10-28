package handlers

import (
	"net/http"
)

func (apiCfg *ApiConfig) HandlerGetSimilarBooksByDesc(w http.ResponseWriter, r *http.Request) {

	type params struct {
		Isbn  int32  `json:"isbn"`
		Title string `json:"title"`
	}

	// Take the isbn, do a search on the embeddings to find the closest ones,

}
