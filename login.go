package main

import (
	"context"
	"fmt"
)

func Login(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	_, err := s.DB.GetUserByName(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	s.Config.CurrentUserName = cmd.Args[0]
	return saveConfig(s.Config)
}
