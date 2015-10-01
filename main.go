package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coralproject/core/database/postgres"
)

func main() {

	//Connection to the DB
	db, err := postgres.Open()

	if err != nil {
		log.Fatal("Could not connect to the database.")
	}
	defer postgres.Close(db)

	http.HandleFunc("/", index)
	http.HandleFunc("/comment/new", newComment)
	http.HandleFunc("/comments", comments)
	http.HandleFunc("/comments/{comment_id}", showComment)

	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	fmt.Fprint(w, "Index")
}

func newComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	fmt.Fprint(w, "New Comment")
}

func comments(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		log.Fatal("Only GET accepted")
		return
	}

	fmt.Fprint(w, "Comments")
}

func showComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		log.Fatal("Only GET accepted")
		return
	}

	fmt.Fprint(w, "Show Comment")
}
