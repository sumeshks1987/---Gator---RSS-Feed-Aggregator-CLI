package main

import (
	"context"
	"fmt"
)

func Feeds(s *State, cmd command) error {
	feeds, err := s.DB.GetFeedsWithUsers(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s (%s)\n", feed.Name, feed.UserName)
	}

	return nil
}
