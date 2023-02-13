package player

import (
	"player-manager/internal/xo"
	"testing"
)

func Test_incrementItem(t *testing.T) {
	tests := []struct {
		name             string
		item             *xo.Item
		quantity         int
		expectedQuantity int
	}{
		{
			name:             "try 0=>1",
			item:             &xo.Item{},
			quantity:         1,
			expectedQuantity: 1,
		},
		{
			name:             "try 0=>0",
			item:             &xo.Item{},
			quantity:         0,
			expectedQuantity: 1,
		},
		{
			name:             "try 0=>-1",
			item:             &xo.Item{},
			quantity:         -1,
			expectedQuantity: 1,
		},
		{
			name:             "try 1=>2",
			item:             &xo.Item{Quantity: 1},
			quantity:         1,
			expectedQuantity: 2,
		},
		{
			name:             "try 1=>11",
			item:             &xo.Item{Quantity: 1},
			quantity:         10,
			expectedQuantity: 11,
		},
		{
			name:             "try 100=>101",
			item:             &xo.Item{Quantity: 100},
			quantity:         1,
			expectedQuantity: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incrementItem(tt.item, tt.quantity)
		})
	}
}

func Test_decrementItem(t *testing.T) {
	tests := []struct {
		name             string
		item             *xo.Item
		quantity         int
		expectedQuantity int
	}{
		{
			name:             "try 0=>-1",
			item:             &xo.Item{Quantity: 0},
			quantity:         1,
			expectedQuantity: 0,
		},
		{
			name:             "try 0=>0",
			item:             &xo.Item{Quantity: 0},
			quantity:         0,
			expectedQuantity: 0,
		},
		{
			name:             "try 2=>1",
			item:             &xo.Item{Quantity: 2},
			quantity:         1,
			expectedQuantity: 1,
		},
		{
			name:             "try 11=>1",
			item:             &xo.Item{Quantity: 11},
			quantity:         10,
			expectedQuantity: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decrementItem(tt.item, tt.quantity)
		})
	}
}
