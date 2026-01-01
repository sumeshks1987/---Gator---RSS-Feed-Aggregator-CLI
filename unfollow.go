package main

import (
	"context"
	"fmt"

	"rss/internal/database"
)

func Unfollow(s *State, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}

	ctx := context.Background()

	feed, err := s.DB.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s has unfollowed %s\n", user.Name, feed.Name)
	return nil
}
