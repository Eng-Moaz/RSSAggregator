package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/Eng-Moaz/RSSAggregator/internal/database"
	"github.com/google/uuid"
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
	parsedTime, err := time.Parse(time.RFC1123Z, post.PubDate)
	if err != nil {
    	    parsedTime, err = time.Parse(time.RFC1123, post.PubDate)
	}

	params := database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: post.Title,
		Url: post.Link,
		Description: sql.NullString{String: post.Description, Valid: post.Description != ""},
		PublishedAt: sql.NullTime{Time: parsedTime, Valid: err == nil,},
		FeedID: feedToFetch.ID,
		}		
	_, err = s.db.CreatePost(context.Background(), params)
	if err != nil{
		fmt.Printf("failed to create post: %v", err)	
		continue
		}
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
	return nil
} 

func HandlerBrowse(s *state, cmd Command) error {
	current_user, err := s.db.GetUser(context.Background(), s.cfg.USERNAME)
	if err != nil{
		return fmt.Errorf("Failed to create field: %v", err)
	}
	var limit int32 = 2
	if len(cmd.Args) == 1 {
    	    parsedLimit, err := strconv.Atoi(cmd.Args[0])
    	    if err != nil {
        	return fmt.Errorf("invalid limit provided: %v", err)
    	    }
    	    limit = int32(parsedLimit)
	}
	params := database.GetPostsForUserParams{
		UserID: current_user.ID,
		Limit: limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	for _, post := range posts{
		fmt.Printf("Post Title: %v\n%v\n", post.Title, post.Description)
	}
	return nil
}
