package types

import (
	"fmt"
)

func CalculateTotalValue(s Storer, ids []ID) (float64, error) {
	var res float64
	for _, id := range ids {
		product, err := s.GetProduct(id)
		if err != nil {
			return 0, fmt.Errorf("Error while calculating total value: %w", err)
		}

		res += float64(product.Quantity) * product.Price
	}

	return res, nil
}
