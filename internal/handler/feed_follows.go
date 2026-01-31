package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
)

func HandlerFollowing(s *config.State, cmd config.Command, user database.User) error {
	feeds, err := s.DB.GetFeedFollowsForUsers(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("The Feeds name: %s has user %s following!\n", feed.FeedName, feed.UserName)
	}

	os.Exit(0)
	return nil
}
