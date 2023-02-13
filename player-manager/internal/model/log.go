package model

import (
	"encoding/json"
	"player-manager/internal/xo"
	"time"
)

type Log struct {
	CreatedAt time.Time `json:"-"`
	Msg       string    `json:"msg"`
}

func (l Log) ToXoLog() *xo.Log {
	b, _ := json.Marshal(l)
	return &xo.Log{
		Msg:       b,
		CreatedAt: time.Now(),
	}
}

func ToLog(xoLog *xo.Log) (log Log) {
	_ = json.Unmarshal(xoLog.Msg, &log)
	log.CreatedAt = xoLog.CreatedAt
	return
}
