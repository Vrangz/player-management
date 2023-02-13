package model

import (
	"player-manager/internal/xo"
)

type Player struct {
	ID       int `json:"-"`
	Username string
}

func (p Player) ToXoPlayer() *xo.Player {
	return &xo.Player{
		ID:       p.ID,
		Username: p.Username,
	}
}

func ToPlayer(p *xo.Player) Player {
	return Player{
		ID:       p.ID,
		Username: p.Username,
	}
}
