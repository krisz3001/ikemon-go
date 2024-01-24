package main

import (
	"net/http"
)

type Signup struct {
	Username string
	Password string
	Email    string
}

type Login struct {
	Username string
	Password string
}

func UsernameExists(name string) bool {
	row := db.QueryRow("SELECT Username FROM `USERS` WHERE Username = ?", name)
	var username string
	err := row.Scan(&username)
	return err == nil
}

func ValidateSignupForm(r *http.Request) (*Signup, []PostError) {
	r.ParseMultipartForm(0)
	errors := make([]PostError, 0)
	data := &Signup{}
	post := r.PostForm

	if !IsSet(&post, "username") {
		errors = append(errors, "Username is not given")
	} else if UsernameExists(post.Get("username")) {
		errors = append(errors, "Username is taken")
	}
	if !IsSet(&post, "password") {
		errors = append(errors, "Password is not given")
	} else if !IsSet(&post, "passwordagain") {
		errors = append(errors, "Confirm password")
	} else if post.Get("password") != post.Get("passwordagain") {
		errors = append(errors, "Passwords do not match")
	}
	if !IsSet(&post, "email") {
		errors = append(errors, "Email is not given")
	}
	if len(errors) == 0 {
		data.Username = post.Get("username")
		data.Password = post.Get("password")
		data.Email = post.Get("email")
		return data, nil
	}
	return nil, errors
}

func ValidateLoginForm(r *http.Request) (*Login, []PostError) {
	r.ParseMultipartForm(0)
	errors := make([]PostError, 0)
	data := &Login{}
	post := r.PostForm

	if !IsSet(&post, "username") {
		errors = append(errors, "Username is not given")
	}
	if !IsSet(&post, "password") {
		errors = append(errors, "Password is not given")
	}
	if len(errors) == 0 {
		data.Username = post.Get("username")
		data.Password = post.Get("password")
		if passwordMatches := ValidatePassword(data); !passwordMatches {
			errors = append(errors, "Invalid credentials")
			return nil, errors
		}
		return data, nil
	}
	return nil, errors
}

func Controller_Signup(w http.ResponseWriter, r *http.Request) {
	session, user := ProcessSession(w, r)
	if user != nil {
		SendErrors(w, nil)
		return
	}
	switch r.Method {
	case http.MethodPost:
		data, errors := ValidateSignupForm(r)
		if len(errors) == 0 {
			err := AddUser(data, session)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		SendErrors(w, &errors)
	}
}

func Controller_Login(w http.ResponseWriter, r *http.Request) {
	session, user := ProcessSession(w, r)
	if user != nil {
		SendErrors(w, nil)
		return
	}
	switch r.Method {
	case http.MethodPost:
		data, errors := ValidateLoginForm(r)
		if len(errors) == 0 {
			err := LoginUser(data, session)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		SendErrors(w, &errors)
	}
}

func Controller_Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := ProcessSession(w, r)
	session.UserId = ""
	session.Page = 0
	session.TypeId = 0
	err := UpdateSession(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
