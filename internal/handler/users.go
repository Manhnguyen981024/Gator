package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
)

func HandlerUsers(s *config.State, cmd config.Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, u := range users {
		if u.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}
