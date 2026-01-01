package web

import (
	"net/http"

	"rss/internal/database"
)

type Server struct {
	DB *database.Queries
}

func New(db *database.Queries) *Server {
	return &Server{DB: db}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/feeds", s.handleFeeds)
	mux.HandleFunc("/posts", s.handlePosts)
	mux.HandleFunc("/login", s.handleLogin)
	mux.HandleFunc("/logout", s.handleLogout)

	return http.ListenAndServe(addr, mux)
}
