package main

import "net/http"

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
