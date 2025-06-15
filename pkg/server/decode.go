package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decode(r *http.Request, into any) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(into); err != nil {
		return fmt.Errorf("could not decode body: %w", err)
	}
	return nil
}
