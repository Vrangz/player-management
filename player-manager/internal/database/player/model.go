package player

import "player-manager/internal/xo"

func incrementItem(item *xo.Item, quantity int) {
	if quantity == 0 {
		quantity = 1
	}

	item.Quantity += quantity
	if item.Quantity > 100 {
		item.Quantity = 100
	}
}

func decrementItem(item *xo.Item, quantity int) {
	if quantity == 0 {
		quantity = 1
	}

	item.Quantity -= quantity
	if item.Quantity < 0 {
		item.Quantity = 0
	}
}
