package main

import (
	"log"

	"github.com/the-developer-guy/LoTo/internals"
	"github.com/the-developer-guy/LoTo/web"
)

func main() {
	config, err := internals.GetConfig()
	if err != nil {
		log.Fatalf("No valid config file found: %v", err)
	}
	web.StartServer(config)
}
