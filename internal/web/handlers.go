package web

import (
	"context"
	"html/template"
	"net/http"

	"rss/internal/database"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/web/templates/layout.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) handleFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"internal/web/templates/layout.html",
		"internal/web/templates/feeds.html",
	))

	tmpl.Execute(w, feeds)
}

func (s *Server) handlePosts(w http.ResponseWriter, r *http.Request) {
	user, err := s.getLoggedInUser(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	posts, err := s.DB.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  25,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"internal/web/templates/layout.html",
		"internal/web/templates/posts.html",
	))

	tmpl.Execute(w, posts)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles(
			"internal/web/templates/layout.html",
			"internal/web/templates/login.html",
		))
		tmpl.Execute(w, nil)
		return
	}

	// POST
	username := r.FormValue("username")

	user, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "gator_user",
		Value: user.ID.String(),
		Path:  "/",
	})

	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}

func (s *Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "gator_user",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (s *Server) getLoggedInUser(r *http.Request) (database.User, error) {
	cookie, err := r.Cookie("gator_user")
	if err != nil {
		return database.User{}, err
	}

	// Convert string → pgtype.UUID
	var id pgtype.UUID
	err = id.Scan(cookie.Value)
	if err != nil {
		return database.User{}, err
	}

	return s.DB.GetUserByID(context.Background(), id)
}
