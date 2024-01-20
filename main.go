package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

const PORT int = 3000

func init() {
	ikemon_db := Database{DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME}
	db = ikemon_db.Connect()
	fmt.Println("\nBackend is ready!")
}

func main() {
	defer db.Close()

	http.HandleFunc("/cards", Controller_Cards)
	http.HandleFunc("/signup", Controller_Signup)
	http.HandleFunc("/login", Controller_Login)
	http.HandleFunc("/logout", Controller_Logout)

	fmt.Println("Listening on port " + fmt.Sprint(PORT) + "...")
	http.ListenAndServe(":"+fmt.Sprint(PORT), nil)

}
