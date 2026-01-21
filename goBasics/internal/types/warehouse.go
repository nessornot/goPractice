package types

import (
	"errors"
)

func NewWarehouse(cap int) *Warehouse {
	return &Warehouse{
		ProductList: make(map[ID]Product, cap),
	}
}


func (w *Warehouse) AddProduct(p Product) error {
	_, inMap := w.ProductList[p.ID]
	if inMap {
		return errors.New("Product is already in product list")
	}

	w.ProductList[p.ID] = p

	return nil
}

func (w *Warehouse) GetProduct(id ID) (Product, error) {
	product, inMap := w.ProductList[id]
	if !inMap {
		return Product{}, errors.New("Product not found")
	}

	return product, nil
}

func (w *Warehouse) UpdateStock(id ID, delta int) error {
	product, inMap := w.ProductList[id]
	if !inMap {
		return errors.New("Product not found")
	}

	product.Quantity += delta

	if product.Quantity < 0 {
		return errors.New("Insufficient stock")
	}

	w.ProductList[id] = product

	return nil
}
