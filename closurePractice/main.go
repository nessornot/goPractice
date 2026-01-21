package main

import (
	"fmt"
	models "closurePractice/internal/models"
)

func main() {
	// slice of orders
	orders := []models.Order{
		{ID: 0, Amount: 300.0},
		{ID: 1, Amount: 600.0},
		{ID: 2, Amount: 900.0},
		{ID: 3, Amount: 1200.0},
		{ID: 4, Amount: 1500.0},
	}

	calc := createDiscountCalculator(0.1)

	newOrders := filterOrders(orders, func(o models.Order) bool {
		return o.Amount > 200
	})

	for _, order := range newOrders {
		fmt.Println(calc(order.Amount))
	}
}
	
func createDiscountCalculator(baseDiscount float64) func(amount float64) float64 {
	return func(amount float64) float64 {
		if amount > 1000 {
			// discount that adds to base if amount > 1000
			addedDiscount := 0.05
			return amount * (1.0 - baseDiscount - addedDiscount)
		}

		return amount * (1.0 - baseDiscount)
	}
}

func filterOrders(orders []models.Order, predicate func(models.Order) bool) []models.Order {
	trueOrders := make([]models.Order, 0, len(orders))

	for _, order := range orders {
		if predicate(order) {
			trueOrders = append(trueOrders, order)
		}
	}

	return trueOrders
}
