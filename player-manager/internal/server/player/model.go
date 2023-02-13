package player

import "player-manager/internal/model"

// swagger:parameters playerInfo playerAddItem playerDeleteItem playerItems build
type PlayerParameter struct {
	// The username of the player
	//
	// in: path
	// required: true
	Username string `json:"username"`
}

// swagger:parameters playerAddItem playerDeleteItem
type ItemParameter struct {
	// Item name
	// in: path
	// required: true
	// enum: wood,stone,food
	Item string `json:"item"`
	// Item quantity
	// in: query
	// description: if not provided then default is 1; if provided quantity is lower than 1, it will be set to default
	Quantity string `json:"quantity"`
}

// swagger:response NoContentResponse
type NoContentResponse struct {
	// No content
}

// swagger:model PlayerResponse
type PlayerResponse struct {
	// username
	// in: body
	Username string `json:"username"`
}

// swagger:model PlayerItemsResponse
type PlayerItemsResponse struct {
	// items
	// in: body
	Items []Item `json:"items"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func ToPlayerResponse(p model.Player) PlayerResponse {
	return PlayerResponse{Username: p.Username}
}

func ToItemsResponse(items model.Items) (pir PlayerItemsResponse) {
	for _, item := range items {
		pir.Items = append(pir.Items, Item{
			Name:     item.Name,
			Quantity: item.Quantity,
		})
	}
	return
}
