package server

import (
	"encoding/json"
	"net/http"

	"github.com/blacksails/rocket-messages/pkg/api"
)

func respond(w http.ResponseWriter, status int, response any) {
	w.WriteHeader(status)
	if response == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		respondErr(w, http.StatusInternalServerError, err)
	}
}

func respondErr(w http.ResponseWriter, status int, err error) {
	respond(w, http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
}
