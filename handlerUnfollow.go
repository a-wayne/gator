package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/a-wayne/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("must provide a feed url to unfollow")
	}

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	})

	if err != nil {
		return fmt.Errorf("cannot unfollow feed: %s", err)
	}
	return nil
}
