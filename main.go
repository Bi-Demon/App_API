package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Credentials have informations [email] & [password] for storing to database
type Credentials struct {
	email    string
	password string
}

//users for login API
type users struct {
	Email    string
	Password string
}

func main() {

	UseDatabase()
	// initEvents()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomeHandler).Methods("GET")

	myRouter.HandleFunc("/login", LoginHandler).Methods("POST")
	myRouter.HandleFunc("/signup", SignupHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":7000", myRouter))
}

//----------------------------------------------------------------------

// HomeHandler link to home page API
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello")

	fmt.Fprintln(w, "SIGN UP OR LOGIN")
}

/*LoginHandler get user's information for Logging in
return result user existed or not */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var Guest Credentials
	Guest.email, Guest.password = r.FormValue("email"), r.FormValue("password")

	// fmt.Printf("Email => [%s]\n", email)
	// fmt.Printf("Password => [%s]\n", password)

	result := FindUser(Guest.email, Guest.password)

	if result == 0 {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	OldUser := users{}

	OldUser.Email = Guest.email
	OldUser.Password = Guest.password

	Myuser, err := json.Marshal(OldUser)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(Myuser)

}

// SignupHandler get user's information for Signing up to API
func SignupHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var newUsers Credentials
	newUsers.email, newUsers.password = r.FormValue("email"), r.FormValue("password")

	// fmt.Printf("Email => [%s]\n", email)
	// fmt.Printf("Password => [%s]\n", password)

	result := ExistUser(newUsers.email)

	if result == 1 {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	AddUser(newUsers.email, newUsers.password)
}

// AddUser make a SQL's Query  to add user's information to database
func AddUser(email, password string) {

	QueryStmt := `
	INSERT INTO users(email, password) 
	VALUES($1,$2)`

	_, err := db.Exec(QueryStmt, email, password)

	if err != nil {
		panic(err)
	}

}

// FindUser use database to find if user's information exsit or not
func FindUser(email, password string) int64 {

	QueryStmt := `
	SELECT * FROM users
	WHERE email=$1 AND password=$2
	`

	result, err := db.Exec(QueryStmt, email, password)

	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rows
}

//ExistUser check in if email Registered
func ExistUser(email string) int64 {

	QueryStmt := `
	SELECT * FROM users
	WHERE email=$1
	`

	result, err := db.Exec(QueryStmt, email)

	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rows
}

// UseDatabase open connection to PostgreSQL
func UseDatabase() {
	var err error
	conStr := "user=postgres password='123456' dbname=server_api sslmode=disable"
	db, err = sql.Open("postgres", conStr)

	err = db.Ping()

	if err != nil {
		log.Fatal("Error: Could not establish connection with the database")
	}

	fmt.Println("Connected to database")

}
