package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset: %s", err)
	}
	fmt.Println("successfully reset")
	return nil
}
