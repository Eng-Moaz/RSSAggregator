package main

import (
	"context"
	"fmt"
	"time"
)

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		return fmt.Errorf("Can't find next feed: %v", err)
	}

	feedToFetch, err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil{
		return fmt.Errorf("Couldn't mark as fetched: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil{
		return fmt.Errorf("Couldn't fetch feed: %v", err)
	}

	for _, post := range rssFeed.Channel.Item{
		fmt.Println(post.Title)
	}

	return nil
}


func HandlerAgg(s *state, cmd Command) error {
	if len(cmd.Args) <= 0{
		return fmt.Errorf("Invalid arguments length")
	}
	timeToParse := cmd.Args[0]
	timeParsed, err := time.ParseDuration(timeToParse)
	if err != nil{
		return fmt.Errorf("Invalid time argument")
	}

	fmt.Printf("Collecting feeds every %v\n", timeParsed)
	ticker := time.NewTicker(timeParsed)
	for ; ; <-ticker.C {
			scrapeFeeds(s)
	}
} 
