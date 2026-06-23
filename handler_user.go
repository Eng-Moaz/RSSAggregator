package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Eng-Moaz/RSSAggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerLogin(s *state, cmd Command) error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Invalid arguments length")
	}
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("username is required")
	}
	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil{
		return fmt.Errorf("user doesn't exist")
	}
	err = s.cfg.SetUser(name)
	if err != nil{
		return fmt.Errorf("Couldn't set the username to %v", name)
	}
	fmt.Printf("Username set to %v", name)
	return nil	
}

func HandlerRegister(s *state, cmd Command) error {
	if len(cmd.Args) == 0{
		return fmt.Errorf("Invalid arguments")
	}
	params := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	}
	_, err := s.db.CreateUser(context.Background(), params)
	if err != nil{
		return fmt.Errorf("Failed to create user")
	}
	s.cfg.SetUser(params.Name)
	fmt.Printf("User was created with an id of %v and name of %v\n", params.ID, params.Name)
	return nil
}

func HandlerReset(s *state, cmd Command) error{
	err := s.db.Reset(context.Background())	
	if err != nil{
		return fmt.Errorf("Failed to reset")
	}
	fmt.Println("Successfully reset the table")
	return nil
}

func HandlerUsers(s *state, cmd Command) error {
	names, err := s.db.ListNames(context.Background())	
	if err != nil{
		return fmt.Errorf("Failed to list names")
	}
	for _, name := range names{
		if name == s.cfg.USERNAME{
			fmt.Printf("* %v (current)\n", name)
		}else{
			fmt.Printf("* %v\n", name)
		}
	}
	return nil
}

func HandlerAgg(s *state, cmd Command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil{
		return fmt.Errorf("Error in fetching feed: %v", err)
	}

	fmt.Println(rss)
	return nil
}

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

	return nil
}
