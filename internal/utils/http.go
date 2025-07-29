package utils

import (
	"encoding/json"
	"net/http"

	t "github.com/MrRainbow0704/DnD/internal/types"
)

func SendJSON(w http.ResponseWriter, s int, m t.AnyMap) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	if m == nil {
		m = t.AnyMap{}
	}

	json.NewEncoder(w).Encode(m)
}

func ErrorJSON(w http.ResponseWriter, s int, e t.ErrorMap) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	m := t.AnyMap{}
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
