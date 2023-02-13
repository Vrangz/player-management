package player

import (
	"context"
	"database/sql"
	"player-manager/internal/database/log"
	"player-manager/internal/database/player"
	"player-manager/internal/model"
)

type Controller struct {
	playerRepository PlayerRepository
}

//go:generate mockery --name PlayerRepository --output ${pwd}/internal/mocks/
type PlayerRepository interface {
	GetPlayer(ctx context.Context, username string) (model.Player, error)
	ListItems(ctx context.Context, username string) (model.Items, error)
	AddItem(ctx context.Context, username string, item string, quantity int) error
	DeleteItem(ctx context.Context, username string, item string, quantity int) error
	Build(ctx context.Context, username string) error
}

func NewController(db *sql.DB) *Controller {
	loggerRepository := log.NewRepository(db)
	playerRepository := player.NewRepository(db, loggerRepository)
	return &Controller{
		playerRepository: playerRepository,
	}
}
