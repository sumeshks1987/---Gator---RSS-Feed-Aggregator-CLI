package main

import (
	"context"
	"fmt"
	"strconv"

	"rss/internal/database"
)

func Browse(s *State, cmd command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = int32(n)
	}

	posts, err := s.DB.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  limit,
		},
	)
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Printf(
			"\n[%s]\n%s\n%s\n",
			p.FeedName,
			p.Title,
			p.Url,
		)
	}

	return nil
}
