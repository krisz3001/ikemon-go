package main

import (
	"net/http"
)

func Controller_Cards(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	_, _ = ProcessSession(w, r)
	switch r.Method {
	case http.MethodGet:
		result, err := GetAllCards()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendJSON(w, result, nil, "cards")
	}
}
