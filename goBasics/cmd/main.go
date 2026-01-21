package main

import (
	"fmt"
	types "miniWarehouseCLI/internal/types"
)


func main() {
	fmt.Println("program started")

	warehouse := types.NewWarehouse(10)

	err := warehouse.AddProduct(types.Product{
		ID:			"1",
		Name:		"Beef",
		Quantity:	5,
		Price:		500.0,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = warehouse.AddProduct(types.Product{
		ID:			"2",
		Name:		"Salad",
		Quantity:	10,
		Price:		200.0,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = warehouse.UpdateStock("1", 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := types.CalculateTotalValue(warehouse, []types.ID{"1", "2"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
	
	fmt.Println("program ended")
}
