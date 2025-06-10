package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/a-wayne/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("register expects a username")
	}

	_, err := s.db.GetUserByName(context.Background(), cmd.args[0])
	if err == nil {
		return errors.New("username already exists")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})

	if err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User '%s' was created: %s", user.Name, user)
	return nil
}
