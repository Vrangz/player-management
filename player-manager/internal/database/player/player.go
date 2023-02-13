package player

import (
	"context"
	"database/sql"
	"fmt"
	"player-manager/internal/model"
	"player-manager/internal/xo"
	"strings"

	"github.com/pkg/errors"
)

type Repository struct {
	db *sql.DB
	l  logger
}

type logger interface {
	Log(ctx context.Context, msg string) error
}

func NewRepository(db *sql.DB, l logger) *Repository {
	return &Repository{db, l}
}

func (r *Repository) GetPlayer(ctx context.Context, username string) (model.Player, error) {
	xoPlayer, err := xo.PlayerByUsername(ctx, r.db, username)
	if err != nil {
		return model.Player{}, errors.Wrap(err, errMsgGetPlayerFailure)
	}

	return model.ToPlayer(xoPlayer), nil
}

func (r *Repository) ListItems(ctx context.Context, username string) (model.Items, error) {
	xoPlayer, err := r.GetPlayer(ctx, username)
	if err != nil {
		return model.Items{}, err
	}

	xoItems, err := xo.ItemsByPlayerID(ctx, r.db, xoPlayer.ID)
	if err != nil {
		return model.Items{}, errors.Wrap(err, errMsgGetItemsFailure)
	}

	return model.ToItems(xoItems), nil
}

func (r *Repository) AddItem(ctx context.Context, username string, item string, quantity int) error {
	updatedItem := strings.ToLower(item)

	if _, ok := model.AvailableItems[updatedItem]; !ok {
		return errors.New(errMsgUnknownItem)
	}

	player, err := r.GetPlayer(ctx, username)
	if err != nil {
		return errors.Wrap(err, errMsgGetPlayerFailure)
	}

	xoItem, err := xo.ItemByPlayerIDName(ctx, r.db, player.ID, updatedItem)
	if err != nil {
		return errors.Wrap(err, errMsgGetItemFailure)
	}

	incrementItem(xoItem, quantity)

	err = xoItem.Update(ctx, r.db)
	if err != nil {
		return errors.Wrap(err, errMsgUpdateItemFailure)
	}

	_ = r.l.Log(ctx, fmt.Sprintf("user '%s' received %d %s", username, quantity, item))

	return nil
}

func (r *Repository) DeleteItem(ctx context.Context, username string, item string, quantity int) error {
	updatedItem := strings.ToLower(item)

	if _, ok := model.AvailableItems[updatedItem]; !ok {
		return errors.New(errMsgUnknownItem)
	}

	player, err := r.GetPlayer(ctx, username)
	if err != nil {
		return errors.Wrap(err, errMsgGetPlayerFailure)
	}

	xoItem, err := xo.ItemByPlayerIDName(ctx, r.db, player.ID, updatedItem)
	if err != nil {
		return errors.Wrap(err, errMsgGetItemFailure)
	}

	decrementItem(xoItem, quantity)

	err = xoItem.Update(ctx, r.db)
	if err != nil {
		return errors.Wrap(err, errMsgUpdateItemFailure)
	}

	_ = r.l.Log(ctx, fmt.Sprintf("user '%s' lost %d %s", username, quantity, item))

	return nil
}

func (r *Repository) Build(ctx context.Context, username string) error {
	player, err := r.GetPlayer(ctx, username)
	if err != nil {
		return errors.Wrap(err, errMsgGetPlayerFailure)
	}

	stone, err := xo.ItemByPlayerIDName(ctx, r.db, player.ID, model.Stone)
	if err != nil {
		return errors.Wrap(err, errMsgGetItemFailure)
	}

	wood, err := xo.ItemByPlayerIDName(ctx, r.db, player.ID, model.Wood)
	if err != nil {
		return errors.Wrap(err, errMsgGetItemFailure)
	}

	if stone.Quantity < 10 || wood.Quantity < 30 {
		return errors.New(errMsgNotEnoughItemsToBuild)
	}

	decrementItem(stone, 10)
	decrementItem(wood, 30)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, errMsgTransactionBeginFailure)
	}

	defer func() { _ = tx.Rollback() }()

	if err = stone.Update(ctx, tx); err != nil {
		return errors.Wrap(err, errMsgUpdateItemFailure)
	}

	if err = wood.Update(ctx, tx); err != nil {
		return errors.Wrap(err, errMsgUpdateItemFailure)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, errMsgTransactionCommitFailure)
	}

	_ = r.l.Log(ctx, fmt.Sprintf("user '%s' used 30 wood and 10 stone to build", username))

	return nil
}

func (r *Repository) ConsumeFood(ctx context.Context) error {
	items, err := xo.ItemsByName(ctx, r.db, model.Food)
	if err != nil {
		return errors.Wrap(err, errMsgGetItemsFailure)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, errMsgTransactionBeginFailure)
	}

	defer func() { _ = tx.Rollback() }()

	for _, item := range items {
		if item.Quantity <= 0 {
			continue
		}

		decrementItem(item, 1)

		if err = item.Update(ctx, tx); err != nil {
			return errors.Wrap(err, errMsgUpdateItemFailure)
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, errMsgTransactionCommitFailure)
	}

	return nil
}
