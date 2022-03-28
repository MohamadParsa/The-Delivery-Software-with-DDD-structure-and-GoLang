// Package aggregates holds aggregates that combines many entities into a full object.
package aggregate

import (
	"domain/customer"
	"domain/product"

	entity "entity"
)

type Order struct {
	Customer customer.Customer
	Pack     *entity.Pack
	Products []*product.Product
	Cost     *entity.Price
}
