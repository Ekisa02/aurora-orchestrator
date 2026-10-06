package main

import (
	"log"
	"os"

	"github.com/example/aurora-orchestrator/internal/agent"
	"github.com/example/aurora-orchestrator/internal/handlers"
	"github.com/example/aurora-orchestrator/internal/router"
	"github.com/example/aurora-orchestrator/internal/store"
)

func main() {
	dsn := os.Getenv("AURORA_DB_DSN")
	if dsn == "" {
		dsn = "aurora.db"
	}

	db, err := store.New(dsn)
	if err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}

	manager := agent.NewManager(db)

	h := handlers.New(manager)
	r := router.New(h)

	addr := ":8080"
	if v := os.Getenv("AURORA_ADDR"); v != "" {
		addr = v
	}

	log.Printf("starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
