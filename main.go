package main

import (
	"log"
	"fmt"
	"github.com/Eng-Moaz/RSSAggregator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil{
		log.Fatal("Couldn't read the json file")
	}
	cfg.SetUser("Moaz")
	cfg, err = config.Read()
	if err != nil{
		log.Fatal("Couldn't read the json file")
	}

	fmt.Println(cfg)
}
