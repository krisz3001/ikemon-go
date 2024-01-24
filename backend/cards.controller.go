package main

import (
	"net/http"
	"strconv"
)

func Controller_Cards(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	_, _ = ProcessSession(w, r)
	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()

		page, pageErr := strconv.Atoi(queryParams.Get("page"))
		typeid, typeErr := strconv.Atoi(queryParams.Get("type"))
		id := queryParams.Get("id")

		var res []Response
		if id != "" {
			card := GetCardById(id)
			if card != nil {
				SendJSON(w, card)
				return
			}
		} else if pageErr != nil || typeErr != nil || page < 0 || typeid < 0 {
			result, err := GetAllCards()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			page_count, err := CountPageOfType(0)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			AddResponse(&res, "cards", &result)
			AddResponse(&res, "pages", page_count)
		} else {
			result, err := GetNthNineByType(page, typeid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			page_count, err := CountPageOfType(typeid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			AddResponse(&res, "cards", &result)
			AddResponse(&res, "pages", page_count)
		}
		WrapResponses(w, &res)
	}
}
