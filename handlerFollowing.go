package main

import (
	"context"
	"fmt"

	"github.com/a-wayne/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' is following these feeds:\n", user.Name)

	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed.Name)
	}

	return nil
}
