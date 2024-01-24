package main

import (
	"net/http"
)

func Controller_Types(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	_, _ = ProcessSession(w, r)
	switch r.Method {
	case http.MethodGet:
		data, err := GetTypes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		var res []Response
		AddResponse(&res, "types", data)
		WrapResponses(w, &res)
	}
}
