package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Users have informations [email] & [password] for storing to database
type Users struct {
	email    string
	password string
}

func main() {

	UseDatabase()

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", HomeHandler)

	myRouter.HandleFunc("/login", LoginHandler).Methods("POST")
	myRouter.HandleFunc("/signup", SignupHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":7000", myRouter))
}

// HomeHandler link to home page API
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello")

	fmt.Fprintln(w, "SIGN UP OR LOGIN")
}

/*LoginHandler get user's information for Logging in
return result user existed or not */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email, password := r.FormValue("email"), r.FormValue("password")

	// fmt.Printf("Email => [%s]\n", email)
	// fmt.Printf("Password => [%s]\n", password)

	result := FindUser(email, password)

	if result == 0 {
		fmt.Fprintln(w, "Email or Password is UNKNOWN")
	} else {
		fmt.Fprintln(w, "Welcome back ", email)
	}

}

// SignupHandler get user's information for Signing up to API
func SignupHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	email, password := r.FormValue("email"), r.FormValue("password")

	// fmt.Printf("Email => [%s]\n", email)
	// fmt.Printf("Password => [%s]\n", password)

	AddUser(email, password)

	fmt.Fprintln(w, "ENJOY !")
}

// AddUser make a SQL's Query  to add user's information to database
func AddUser(email, password string) {

	QueryStmt := `
	INSERT INTO "Users"(email, password) 
	VALUES($1,$2)`

	_, err := db.Exec(QueryStmt, email, password)

	if err != nil {
		panic(err)
	}
}

// FindUser use database to find if user's information exsit or not
func FindUser(email, password string) int64 {

	QueryStmt := `
	SELECT * FROM "Users"
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
