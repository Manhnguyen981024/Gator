package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Manhnguyen981024/blog-aggregator/internal/api"
	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s *config.State, cmd config.Command, user database.User) error {
	url := "https://www.wagslane.dev/index.xml"
	fmt.Println("handle agg")
	_, err := api.FetchFeed(context.Background(), url)
	if err != nil {
		fmt.Printf("loi ne %v", err)
		return err
	}

	log.Printf("Collecting feeds every %s...", "2s")

	ticker := time.NewTicker(time.Second * 2)
	for ; ; <-ticker.C {
		scrapeFeeds(s, user)
	}
}

func scrapeFeeds(s *config.State, user database.User) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background(), user.ID)
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)

		return err
	}
	log.Println("Found a feed to fetch!")

	validNullTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	updatedFeed := database.MarkfeedFetchedParams{
		LastFetchedAt: validNullTime,
		UpdatedAt:     validNullTime,
		ID:            feed.ID,
	}

	updErr := s.DB.MarkfeedFetched(context.Background(), updatedFeed)
	if updErr != nil {
		return updErr
	}

	val, err := api.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)

		return err
	}

	insertPostForUser(s, val, feed.ID)

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(val.Channel.Item))

	return nil
}

func insertPostForUser(s *config.State, feeds *api.RSSFeed, feedId uuid.UUID) error {
	for _, v := range feeds.Channel.Item {
		createdPostParam := database.CreatePostParams{
			CreatedAt:   time.Now(),
			Title:       v.Title,
			Url:         v.Link,
			Description: v.Description,
			FeedID:      feedId,
			PublishedAt: parseStringToTime(v.PubDate),
		}
		log.Println(parseStringToTime(v.PubDate))
		createdPost, err := s.DB.CreatePost(context.Background(), createdPostParam)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("We have already fetched the title %s \n", createdPost.Title)
	}
	return nil
}

func parseStringToTime(timeStr string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Printf("error to parse time: ", err)
	}
	return t
}
