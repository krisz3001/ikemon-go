package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Session struct {
	Id      string
	UserId  string
	Page    int
	TypeId  int
	Created string
}

func (s Session) String() string {
	return fmt.Sprintf("\t[Session] ID: %s, UserId: %s, Created: %s", s.Id, s.UserId, s.Created)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetSession(id string) (*Session, error) {
	row := db.QueryRow("SELECT * FROM `SESSIONS` WHERE SessionId = ?", id)

	s := &Session{}

	err := row.Scan(&s.Id, &s.UserId, &s.Page, &s.TypeId, &s.Created)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func CreateSession(w http.ResponseWriter) *Session {
	id := NewSessionId()
	http.SetCookie(w, &http.Cookie{Name: "id", Value: id})
	session := &Session{id, "", 0, 0, time.Now().Format("2006-01-02 15:04:05")}
	_, err := db.Exec("INSERT INTO `SESSIONS` VALUES (?,?,?,?,?)", session.Id, session.UserId, session.Page, session.TypeId, session.Created)
	if err != nil {
		log.Fatal(err.Error())
	}
	return session
}

func UpdateSession(s *Session) error {
	_, err := db.Exec("UPDATE `SESSIONS` SET UserId = ?, Page = ?, TypeId = ? WHERE SessionId = ?", s.UserId, s.Page, s.TypeId, s.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewSessionId() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ProcessSession(w http.ResponseWriter, r *http.Request) (*Session, *User) {
	cookie, err := r.Cookie("id")
	var id string
	var session *Session
	var user *User
	if err == http.ErrNoCookie {
		session = CreateSession(w)
	} else {
		id = cookie.Value
		session, err = GetSession(id)
		if err != nil {
			session = CreateSession(w)
		}
	}
	if session.UserId != "" {
		user, err = GetUserById(session.UserId)
		if err != nil {
			return nil, nil
		}
		return session, user
	}
	return session, nil
}
