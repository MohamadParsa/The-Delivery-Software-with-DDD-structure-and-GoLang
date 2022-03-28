package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrProductNotFound is returned when a product is not found.
	ErrProductNotFound = errors.New("the product was not found in the repository")
	// ErrFailedToAddProduct is returned when the product could not be added to the repository.
	ErrFailedToAddProduct = errors.New("failed to add the product to the repository")
	// ErrUpdateProduct is returned when the product could not be updated in the repository.
	ErrUpdateProduct = errors.New("failed to update the product in the repository")
)
//ProductRepository interface used to replace any repository to store product data.
type ProductRepository interface {
	GetAll() (map[uuid.UUID]*Product, error)
	Get(uuid.UUID) (*Product, error)
	Add(Product) error
	Update(Product) error
}
