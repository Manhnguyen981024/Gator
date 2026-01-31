package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
)

func HandlerUnfollow(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Arguments) < 1 {
		return errors.New("the HandlerAddFeed handler expects a two argument")
	}

	feed, err := s.DB.GetFeedByUrl(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}
	newFeed := database.DeleteFeedFollowByURLParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	errDel := s.DB.DeleteFeedFollowByURL(context.Background(), newFeed)
	if errDel != nil {
		log.Fatalf("Cannot delete new feed %v", err)
		os.Exit(1)
	}

	fmt.Printf("Delete RssFeed successfully!! ")

	return nil
}
