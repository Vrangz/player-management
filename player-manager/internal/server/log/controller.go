package log

import (
	"context"
	"database/sql"
	"player-manager/internal/database/log"
	"player-manager/internal/model"
)

type Controller struct {
	logRepository LogRepository
}

type LogRepository interface {
	GetLogs(ctx context.Context) ([]model.Log, error)
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		logRepository: log.NewRepository(db),
	}
}
