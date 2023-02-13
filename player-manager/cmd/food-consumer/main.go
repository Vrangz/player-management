package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"player-manager/internal/config"
	"player-manager/internal/database"
	"player-manager/internal/database/player"
	"syscall"
	"time"
)

type noopLogger struct{}

func (nl *noopLogger) Log(ctx context.Context, msg string) error {
	return nil
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("config could not be loaded. ", err)
	}

	db, err := database.EstablishConnection(cfg)
	if err != nil {
		log.Fatal("failed to connect to the database. ", err)
	}

	fc := player.NewRepository(db, &noopLogger{})

	ticker := time.NewTicker(time.Minute)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-ticker.C:
			if err = consumeFood(db, fc); err != nil {
				log.Println(err)
			}
		case <-sigCh:
			log.Println("Graceful shutdown...")
			ticker.Stop()
			return
		}
	}
}

type foodConsumer interface {
	ConsumeFood(ctx context.Context) error
}

func consumeFood(db *sql.DB, fc foodConsumer) error {
	return fc.ConsumeFood(context.Background())
}
