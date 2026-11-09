package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)

	if err != nil {
		return err
	}

	fmt.Printf("the user has been set to %s\n", username)
	return nil
}
