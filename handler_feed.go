package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Eng-Moaz/RSSAggregator/internal/database"
	"github.com/google/uuid"
)


func HandlerAddFeed(s *state, cmd Command) error {
	if len(cmd.Args) <= 1 {
		return fmt.Errorf("invalid arguments length")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	current_user, err := s.db.GetUser(context.Background(), s.cfg.USERNAME)
	if err != nil{
		return fmt.Errorf("Failed to create field: %v", err)
	}

	params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: current_user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil{
		return fmt.Errorf("Failed to create field: %v", err)
	}
	
	fmt.Println(feed)

	paramsFollow := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: current_user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), paramsFollow)
	if err != nil{
		return fmt.Errorf("Failed to create follow: %v", err)
	}
	return nil
}

func HandlerListFeeds(s *state, cmd Command) error {
	feedList, err := s.db.ListFeeds(context.Background())
	if err != nil{
		return fmt.Errorf("Failed to get feeds list: %v", err)
	}	

	for _, feed := range feedList{
		name, err := s.db.GetUsername(context.Background(), feed.UserID)
		if err != nil{
			return fmt.Errorf("failed to get name of the user: %v", err)
		}
		fmt.Printf("%v posted a feed with the title %v. Link: %v\n", name, feed.Name, feed.Url)
	}
	return nil
}


