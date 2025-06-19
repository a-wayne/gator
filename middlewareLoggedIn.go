package main

import (
	"context"

	"github.com/a-wayne/gator/internal/database"
)

func middleWareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
