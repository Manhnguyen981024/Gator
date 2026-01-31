package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	fmt.Println("Open request !!")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		log.Fatalf("cannot open a http request to %s", feedURL)
		return nil, err
	}
	fmt.Println("Open request successful!!")

	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Call request fail!!")
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var rssData = RSSFeed{}
	if err := xml.Unmarshal(body, &rssData); err != nil {
		return &rssData, err
	}
	rssData.Channel.Title = html.UnescapeString(rssData.Channel.Title)
	rssData.Channel.Description = html.UnescapeString(rssData.Channel.Description)
	for i, item := range rssData.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssData.Channel.Item[i] = item
	}
	return &rssData, nil
}
