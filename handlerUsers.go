package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %s", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
