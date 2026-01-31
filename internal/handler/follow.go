package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Arguments) < 1 {
		return errors.New("the HandlerAddFeed handler expects a two argument")
	}

	feedUrl := cmd.Arguments[0]

	currentFeed, err := s.DB.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		log.Fatalf("Cannot get feed by url %v", err)
		os.Exit(1)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    currentFeed.ID,
		UserID:    user.ID,
	}

	createdFeed, err := s.DB.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		log.Fatalf("Cannot create new feed follow %v", err)
		os.Exit(1)
	}

	fmt.Printf("The user %s has registered the feed %s successfully!! \n", createdFeed.UserName, createdFeed.FeedName)

	return nil
}
