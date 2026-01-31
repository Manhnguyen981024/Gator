package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Arguments) < 2 {
		return errors.New("the HandlerAddFeed handler expects a two argument")
	}

	dt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: dt,
		UpdatedAt: dt,
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	createdFeed, err := s.DB.CreateFeed(context.Background(), newFeed)
	if err != nil {
		log.Fatalf("Cannot create new feed %v", err)
		os.Exit(1)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    createdFeed.ID,
	}
	_, errFeed := s.DB.CreateFeedFollow(context.Background(), newFeedFollow)
	if errFeed != nil {
		log.Fatalf("Cannot create new feed follow %v", err)
		os.Exit(1)
	}

	fmt.Printf("Create RssFeed %v successfully!! ", createdFeed)

	return nil
}
