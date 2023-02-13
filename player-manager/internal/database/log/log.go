package log

import (
	"context"
	"database/sql"
	"player-manager/internal/model"
	"player-manager/internal/xo"

	"github.com/pkg/errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Log(ctx context.Context, msg string) error {
	return model.Log{
		Msg: msg,
	}.ToXoLog().Insert(ctx, r.db)
}

func (r *Repository) GetLogs(ctx context.Context) ([]model.Log, error) {
	rows, err := r.db.Query("SELECT id, created_at, msg FROM logs ORDER BY created_at DESC")
	if err != nil {
		return []model.Log{}, errors.Wrap(err, errMsgQueryFailure)
	}

	defer rows.Close()

	var logs []model.Log
	var xoLog xo.Log
	for rows.Next() {
		if err = rows.Scan(&xoLog.ID, &xoLog.CreatedAt, &xoLog.Msg); err != nil {
			return []model.Log{}, errors.Wrap(err, errMsgScanFailure)
		}

		logs = append(logs, model.ToLog(&xoLog))
	}

	return logs, nil
}
