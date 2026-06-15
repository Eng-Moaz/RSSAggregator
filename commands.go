package main

import "fmt"

type Command struct{
	Name string
	Args []string
}

type Commands struct{
	commands map[string]func(*state, Command) error	
}

func (c *Commands) run (s *state, cmd Command) error{
	fc, exists := c.commands[cmd.Name]
	if !exists{
		return fmt.Errorf("Function doesn't exist")
	}
	err := fc(s, cmd)
	if err != nil{
		return fmt.Errorf("Couldn't apply the function")
	}
	return nil
}

func (c *Commands) register(name string, f func(*state, Command) error){
	c.commands[name] = f
}


