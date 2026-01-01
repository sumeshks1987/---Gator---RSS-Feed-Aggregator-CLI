package main

import (
	"context"
	"fmt"
	"rss/internal/database"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func scrapeFeeds(s *State) error {
	ctx := context.Background()

	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Fetching:", feed.Name)

	rss, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rss.Channel.Item {
		var published pgtype.Timestamp
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			published = pgtype.Timestamp{Time: t, Valid: true}
		}

		_, err := s.DB.CreatePost(ctx, database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: pgtype.Text{String: item.Description, Valid: true},
			PublishedAt: published,
			FeedID:      feed.ID,
		})

		// Ignore duplicate posts
		if err != nil {
			continue
		}
	}

	return s.DB.MarkFeedFetched(ctx, feed.ID)
}
