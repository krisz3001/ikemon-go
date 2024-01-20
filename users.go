package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   string  `json:"userid"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Money    float64 `json:"money"`
	Admin    bool    `json:"admin"`
	Created  string  `json:"created"`
}

func (u User) String() string {
	return fmt.Sprintf("\t[User] %s, %.2f, %t", u.Username, u.Money, u.Admin)
}

const DEFAULT_MONEY float64 = 1000

func AddUser(s *Signup, session *Session) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), 8)
	if err != nil {
		return err
	}
	userId := uuid.New().String()
	var user *User = &User{userId, s.Username, string(encryptedPassword), s.Email, DEFAULT_MONEY, false, time.Now().Format("2006-01-02 15:04:05")}
	_, err = db.Exec("INSERT INTO `USERS` VALUES (?,?,?,?,?,?,?)", user.UserId, user.Username, user.Password, user.Email, user.Money, user.Admin, user.Created)
	if err != nil {
		return err
	}
	session.UserId = userId
	err = UpdateSession(session)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(l *Login, session *Session) error {
	user, err := GetUserByName(l.Username)
	if err != nil {
		return err
	}
	session.UserId = user.UserId
	err = UpdateSession(session)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(username string) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT * FROM `USERS` WHERE `Username` = ?", username)
	err := row.Scan(&user.UserId, &user.Username, &user.Password, &user.Email, &user.Money, &user.Admin, &user.Created)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id string) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT * FROM `USERS` WHERE `UserId` = ?", id)
	err := row.Scan(&user.UserId, &user.Username, &user.Password, &user.Email, &user.Money, &user.Admin, &user.Created)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ValidatePassword(l *Login) bool {
	row := db.QueryRow("SELECT `Password` FROM `USERS` WHERE `Username` = ?", l.Username)

	var password string
	err := row.Scan(&password)
	if err != nil {
		return false
	}
	matches := bcrypt.CompareHashAndPassword([]byte(password), []byte(l.Password))
	return matches == nil
}
