package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/a-wayne/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		i, err := strconv.ParseInt(cmd.args[0], 10, 16)
		if err != nil {
			return fmt.Errorf("could not parse result limit: %s", err)
		}

		limit = int(i)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("could not get posts for user: %s", err)
	}

	fmt.Printf("Posts for %s\n\n", user.Name)
	for _, post := range posts {
		fmt.Printf("* %s (%s)\n", post.Title, post.Name)
		fmt.Printf("  %s\n", post.Description)
		fmt.Printf("  %s\n\n", post.PublishedAt.String())
		fmt.Printf("  %s\n\n\n", post.Url)
	}
	return nil
}
