package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
)

func HandlerFeeds(s *config.State, cmd config.Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("The feed name: %s with URL: %s of the USER: %s\n", feed.Name, feed.Url, feed.UserName)
	}

	os.Exit(0)
	return nil
}
