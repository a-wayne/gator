package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login expects a username")
	}

	// Verify user
	user, err := s.db.GetUserByName(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("user '%s' does not exist", cmd.args[0])
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error logging in: %s", err)
	}

	fmt.Printf("User '%s' has been logged in.\n", cmd.args[0])
	return nil
}
