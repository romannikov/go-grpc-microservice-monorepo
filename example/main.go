package main

import (
	"github.com/romannikov/example/config"
	"github.com/romannikov/example/server"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Unable to get a config: %v", err)
	}
	if err = server.Start(cfg); err != nil {
		log.Fatalf("Unable to start a server: %v", err)
	}
}
