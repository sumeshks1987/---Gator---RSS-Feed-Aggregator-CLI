package main

import (
	"context"

	"rss/internal/database"
)

type authedHandler func(s *State, cmd command, user database.User) error

func middlewareLoggedIn(handler authedHandler) func(*State, command) error {
	return func(s *State, cmd command) error {
		user, err := s.DB.GetUserByName(
			context.Background(),
			s.Config.CurrentUserName,
		)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
