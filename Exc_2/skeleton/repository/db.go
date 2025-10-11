package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	// drinks := ...
	drinks := []model.Drink{
		{ID: 1, Name: "cola", Price: 5, Description: "Yummy fuzzy drink"},
		{ID: 2, Name: "beer", Price: 3, Description: "Yucky"},
		{ID: 3, Name: "coffee", Price: 4, Description: "Suspicious brown brew"},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Date(2025, time.January, 2, 15, 4, 5, 0, time.UTC), Amount: 10},
		{DrinkID: 2, CreatedAt: time.Date(2025, time.January, 5, 17, 2, 50, 0, time.UTC), Amount: 2},
		{DrinkID: 3, CreatedAt: time.Date(2025, time.January, 3, 10, 40, 25, 0, time.UTC), Amount: 5},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, o := range db.orders {
		totalledOrders[o.DrinkID] += uint64(o.Amount) // because Amount is actually of a different data type, we convert here
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
