package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Eng-Moaz/RSSAggregator/internal/config"
	"github.com/Eng-Moaz/RSSAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct{
	cfg *config.Config
	db *database.Queries
}

func main(){
	cfg, err := config.Read()
	if err != nil{
		log.Fatalf("Couldn't read the config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil{
		log.Fatalf("Couldn't open the database")
	}
	dbQueries := database.New(db)

	programState := state{
		cfg : &cfg,
		db : dbQueries,
	}
	cmds := Commands{
		commands: make(map[string]func(*state, Command) error),
	}
	cmds.register("login", HandlerLogin)
	cmds.register("register", handlerRegister)
	
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
