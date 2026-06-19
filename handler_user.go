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

func handlerRegister(s *state, cmd Command) error {
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
