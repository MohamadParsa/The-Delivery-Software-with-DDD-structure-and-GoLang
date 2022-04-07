package order

import (
	"entity"

	"github.com/google/uuid"
)

type OrderedProduct struct {
	ProductID uuid.UUID
	Count     int
	Price     entity.Price
}
