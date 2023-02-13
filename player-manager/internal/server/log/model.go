package log

import (
	"player-manager/internal/model"
	"time"
)

// swagger:model LogsResponse
type LogsResponse struct {
	// logs
	// in: body
	Logs []Log `json:"items"`
}

type Log struct {
	CreatedAt time.Time `json:"created_at"`
	Msg       string    `json:"msg"`
}

func ToLogsResponse(logs []model.Log) (lr LogsResponse) {
	for _, log := range logs {
		lr.Logs = append(lr.Logs, Log{
			CreatedAt: log.CreatedAt,
			Msg:       log.Msg,
		})
	}
	return
}
