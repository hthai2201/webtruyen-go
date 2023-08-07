package main

import (
	"log"

	"github.com/hthai2201/webtruyen-go/config"
	"github.com/hthai2201/webtruyen-go/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
