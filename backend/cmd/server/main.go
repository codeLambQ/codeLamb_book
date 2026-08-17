package main

import (
	"log"

	"github.com/codeLambQ/codeLamb_book/backend/internal/config"
	"github.com/codeLambQ/codeLamb_book/backend/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New()
	log.Printf("server listening on %s", cfg.Server.Address)
	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
