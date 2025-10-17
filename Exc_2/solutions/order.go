package model

import "time"

type Order struct {
	DrinkID uint64 `json:"drink_id"` // foreign key
	// todo Add fields: CreatedAt (time.Time), Amount with suitable types
	CreatedAt time.Time `json:"created_at"`
	Amount    float32   `json:"amount"`
	// todo json attributes need to be snakecase, i.e. name, created_at, my_variable, ..
}
