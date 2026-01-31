package handler

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
)

func HandlerLogin(s *config.State, cmd config.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	username := cmd.Arguments[0]
	u, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s.Config.SetUser(u.Name)
	fmt.Printf("Login to %s successfully!", u.Name)

	return nil
}
