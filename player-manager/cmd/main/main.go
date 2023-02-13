package main

import (
	"log"

	"player-manager/internal/config"
	"player-manager/internal/database"
	"player-manager/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("config could not be loaded. ", err)
	}

	db, err := database.EstablishConnection(cfg)
	if err != nil {
		log.Fatal("failed to connect to the database. ", err)
	}

	s := server.New(cfg, db)
	log.Fatal(s.Start())
}
