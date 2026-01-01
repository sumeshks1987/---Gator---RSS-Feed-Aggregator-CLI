package main

import (
	"context"
	"fmt"
	"rss/internal/database"
)

func Following(s *State, cmd command, user database.User) error {
	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range follows {
		fmt.Println(f.FeedName)
	}

	return nil
}
