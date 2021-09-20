package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Content string `json:"content"`
}

type Db struct {
	Posts []Post `json:"posts"`
}

type Server struct {
	db Db
}

// RegisterRoutes is a func attached to Server class, s is one Server's variable
// R -- upcase means is a public func, lowercase means is a private func
func (s *Server) RegisterRoutes() {
	http.HandleFunc("/memo/create", s.CreateMemo)
	http.HandleFunc("/memo/list", s.ListMemos)
}

func (s *Server) CreateMemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", 404)
		// return will stop the func
		return
	}

	var post Post
	// := means new create variable
	// NewDecoder will read your json body string, Decode will turn the json body to post type
	// &post is a pointer to the post value
	// if successed, the post will get the value, if valued, it gave you the err
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Invalid memo: "+err.Error(), 400)
		return
	}

	s.db.Posts = append(s.db.Posts, post)

	fmt.Fprintf(w, "Memo created")
}

func (s *Server) ListMemos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", 404)
		return
	}

	err := json.NewEncoder(w).Encode(s.db.Posts)

	if err != nil {
		http.Error(w, "Unable to list memos", 500)
		return
	}
}

func (s *Server) Run() {
	s.RegisterRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		// log a error message and kill the program
		log.Fatal(err)
	}
}

func main() {
	server := Server{
		// give a empty {} when you create a new struct instance
		db: Db{},
	}

	server.Run()
}
