package main

import (
	"log"
	"net/http"

	"github.com/codeLambQ/codeLamb_book/backend/internal/config"
	"github.com/codeLambQ/codeLamb_book/backend/internal/router"
)

func main() {
	cfg := config.Load()

	srv := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: router.New(),
	}

	log.Printf("server listening on %s", cfg.Server.Address)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
