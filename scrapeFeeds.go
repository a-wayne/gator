package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/a-wayne/gator/internal/database"
	"github.com/a-wayne/gator/internal/rss"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed to fetch: %s", err)
	}

	fmt.Printf("scraping feed: %s\n", feed.Name)

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return fmt.Errorf("could not mark feed (%d) as fetched: %s", feed.ID, err)
	}

	rFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %s", err)
	}

	fmt.Printf("feed has %d posts", len(rFeed.Channel.Item))
	for _, x := range rFeed.Channel.Item {
		pub, err := parseTime(x.PubDate)
		if err != nil {
			return fmt.Errorf("could not save post; could not parse pubDate: %s", x.PubDate)
		}

		// Verify this post has not already been saved
		_, err = s.db.GetPostByUrl(context.Background(), x.Link)
		if err == nil {
			//fmt.Printf("Post already saved, skipping (%s)\n", x.Title)
			continue
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       x.Title,
			Url:         x.Link,
			Description: x.Description,
			PublishedAt: pub,
			FeedID:      feed.ID,
		})
		if err != nil {
			return fmt.Errorf("could not save post: %s", err)
		}
	}

	return nil
}

func parseTime(s string) (time.Time, error) {
	layouts := []string{
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("no matching layout for: %s", s)
}
