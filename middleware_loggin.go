package main

import (
	"context"
	"log"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *config.State, cmd config.Command, user database.User) error) func(*config.State, config.Command) error {
	return func(s1 *config.State, c config.Command) error {
		loginUser, err := s1.DB.GetUserByName(context.Background(), s1.Config.CurrentUserName)
		if err != nil {
			log.Fatalf("Cannot get user login %v", err)
			os.Exit(1)
		}
		return handler(s1, c, loginUser)
	}
}
