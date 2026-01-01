package main

import (
	"context"
	"fmt"
)

func Register(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: register <username>")
	}

	username := cmd.Args[0]

	_, err := s.DB.CreateUser(context.Background(), username)
	if err != nil {
		return err
	}

	s.Config.CurrentUserName = username
	return saveConfig(s.Config)
}
