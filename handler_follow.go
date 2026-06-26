package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Eng-Moaz/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *state, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("invalid arguments length")
	}
	url := cmd.Args[0]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.USERNAME)
	if err != nil{
		return fmt.Errorf("Failed to get user: %v", err)
	}

	current_feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
    	    return fmt.Errorf("Failed to get feed by url '%s': %v", url, err)
	}

	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: current_user.ID,
		FeedID: current_feed.ID,
	}
	follows, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil{
		return fmt.Errorf("Failed to create follow: %v", err)
	}
	
	for _, follow := range follows{
		fmt.Printf("%v followed feed {%v}", follow.UserName, follow.FeedName)
	}

	return nil
}

func HandlerFollowing (s *state, cmd Command) error {
	current_user, err := s.db.GetUser(context.Background(), s.cfg.USERNAME)
	if err != nil{
		return fmt.Errorf("Failed to get user: %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), current_user.ID)
	if err != nil{
		return fmt.Errorf("Failed to get feeds: %v", err)
	}

	for _, feed := range feeds{
		feedName, _ := s.db.GetFeedById(context.Background(), feed.FeedID)
		fmt.Println(feedName)
	}
	return nil
}

func HandlerUnfollow(s *state, cmd Command) error {
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("Invalid arguments length")
	}

	current_user, err := s.db.GetUser(context.Background(), s.cfg.USERNAME)
	if err != nil{
		return fmt.Errorf("Failed to get user: %v", err)
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil{
		return fmt.Errorf("Failed to get feed ID: %v", err)
	}

	params := database.DeleteFollowParams{
		UserID : current_user.ID,
		FeedID : feed.ID,
	}
	err = s.db.DeleteFollow(context.Background(), params)
	if err != nil{
		return fmt.Errorf("Failed to DeleteFollow: %v", err)
	}

	return nil
	
}
