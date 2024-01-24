package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	username string
	password string
	host     string
	port     string
	name     string
}

func (d Database) Connect() *sql.DB {
	db, err := sql.Open("mysql", d.username+":"+d.password+"@tcp("+d.host+":"+d.port+")/"+d.name)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Connection failed!")
	} else {
		fmt.Println("Connected to database: " + d.name)
	}
	return db
}
