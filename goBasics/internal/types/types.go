package types

type ID string

type Product struct {
	ID			ID
	Name		string
	Price		float64
	Quantity	int
}

type Warehouse struct {
	ProductList	map[ID]Product
}

type Storer interface {
	AddProduct(p Product) error
	GetProduct(id ID) (Product, error)
	UpdateStock(id ID, delta int) error
}

