package main

import(
	"fmt"
)

func HandlerLogin(s *state, cmd Command) error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Invalid arguments length")
	}
	if len(cmd.Args) <= 0 {
		return fmt.Errorf("username is required")
	}
	name := cmd.Args[0]
	err := s.cfg.SetUser(name)
	if err != nil{
		return fmt.Errorf("Couldn't set the username to %v", name)
	}
	fmt.Printf("Username set to %v", name)
	return nil	
}

