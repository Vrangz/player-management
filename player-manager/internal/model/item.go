package model

import "player-manager/internal/xo"

const (
	Wood  = "wood"
	Stone = "stone"
	Food  = "food"
)

var AvailableItems = map[string]struct{}{
	Wood:  {},
	Stone: {},
	Food:  {},
}

type Items []Item

type Item struct {
	Name     string
	Quantity int
}

func ToItems(xoItems []*xo.Item) Items {
	var items = make(Items, 0, len(xoItems))
	for _, xoItem := range xoItems {
		items = append(items, ToItem(xoItem))
	}
	return items
}

func ToItem(xoItem *xo.Item) Item {
	return Item{
		Name:     xoItem.Name,
		Quantity: xoItem.Quantity,
	}
}
