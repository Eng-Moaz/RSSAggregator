package main

import (
	"log"
	"os"

	"github.com/Eng-Moaz/RSSAggregator/internal/config"
)

type state struct{
	cfg *config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil{
		log.Fatalf("Couldn't read the config: %v", err)
	}
	programState := state{
		cfg : &cfg,
	}
	cmds := Commands{
		commands: make(map[string]func(*state, Command) error),
	}
	cmds.register("login", HandlerLogin)
	
	if len(os.Args) < 2 {
		log.Fatal("Error in arguments length")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(&programState, Command{Name: cmdName, Args: cmdArgs,})
	if err != nil{
		log.Fatalf("error in running: %v", err)
	}

}
