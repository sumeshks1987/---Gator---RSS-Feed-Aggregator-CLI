package main

import (
	"context"
	"fmt"

	"rss/internal/database"
)

func AddFeed(s *State, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	ctx := context.Background()

	feed, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
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

	fmt.Println("Feed added successfully:")
	fmt.Println("ID:", feed.ID)
	fmt.Println("Name:", feed.Name)
	fmt.Println("URL:", feed.Url)

	return nil
}
