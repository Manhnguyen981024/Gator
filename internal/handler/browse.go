package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
)

func HanlderBrowse(s *config.State, cmd config.Command, user database.User) error {
	if len(cmd.Arguments) < 1 {
		return errors.New("the HandlerAddFeed handler expects a 1 argument")
	}

	limit, err := strconv.Atoi(cmd.Arguments[0])

	if err != nil {
		log.Fatalf("please enter a number intead of %s", cmd.Arguments[0])
	}

	params := database.GetPostsByUserIdParams{
		ID:    user.ID,
		Limit: int32(limit),
	}

	posts, err := s.DB.GetPostsByUserId(context.Background(), params)
	if err != nil {
		log.Fatalf("Cannot create new feed %v", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
