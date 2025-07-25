package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, s int, m map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	if m == nil {
		m = map[string]any{}
	}

	json.NewEncoder(w).Encode(m)
}

func ErrorJSON(w http.ResponseWriter, s int, e map[string]error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	m := map[string]any{}
	if e != nil {
		errs := make(map[string]string, len(e))
		for k, v := range e {
			errs[k] = v.Error()
		}
		m["errors"] = errs
	} else {
		m["errors"] = map[string]string{}
	}

	json.NewEncoder(w).Encode(m)
}
