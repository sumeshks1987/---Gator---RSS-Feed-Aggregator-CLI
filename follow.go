package main

import (
	"context"
	"fmt"
	"rss/internal/database"
)

func Follow(s *State, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	ctx := context.Background()

	feed, err := s.DB.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s is now following %s\n", user.Name, feed.Name)
	return nil
}
