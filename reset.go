package main

import "context"

func Reset(s *State, cmd command) error {
	return s.DB.Reset(context.Background())
}
