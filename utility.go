package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Toast struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
	Img  string `json:"img"`
}

type PostError string

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func NewError(errors *[]string, error string) {
	*errors = append(*errors, error)
}

type Response struct {
	Wrapper string
	Obj     any
}

func AddResponse(list *[]Response, wrapper string, obj any) {
	*list = append(*list, Response{Wrapper: wrapper, Obj: obj})
}

func WrapResponses(w http.ResponseWriter, r *[]Response) {
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND_URL)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	var response []byte
	response = append(response, []byte("{")...)
	for i, item := range *r {
		data, err := json.Marshal(item.Obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data = append([]byte("\""+item.Wrapper+"\":"), data...)
		if i != len(*r)-1 {
			data = append(data, []byte(",")...)
		}
		response = append(response, data...)
	}
	response = append(response, []byte("}")...)
	w.Write(response)
}

func SendJSON(w http.ResponseWriter, obj any, wrapper ...string) {
	data, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}

	w.Header().Set("Access-Control-Allow-Origin", FRONTEND_URL)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func SendErrors(w http.ResponseWriter, errors *[]PostError) {
	data, err := json.Marshal(errors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data = append([]byte("{\"errors\":"), data...)
	data = append(data, []byte("}")...)
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND_URL)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func IsSet(post *url.Values, field string) bool {
	return !(!post.Has(field) || Trim(post.Get(field)) == "")
}

func LogRequest(r *http.Request) {
	fmt.Println("[" + time.Now().Format(time.DateTime) + "] " + r.Method + " on " + r.URL.String() + " from " + r.RemoteAddr)
}
