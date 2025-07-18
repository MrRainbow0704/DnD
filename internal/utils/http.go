package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, s int, m map[string]any, e map[string]error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	if e != nil {
		m["errors"] = e
	}
	json.NewEncoder(w).Encode(m)
}