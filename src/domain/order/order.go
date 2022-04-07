package order

//TODO: manage add/remove products.
//TODO: manage count conditions for products.
import (
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	id         uuid.UUID
	customerID uuid.UUID
	products   []*OrderedProduct
}

var (
	//ErrEmptyCustomerID is returned when a customer id is empty.
	ErrEmptyCustomerID = errors.New("customer id is required")
	//ErrEmptyOrderID is returned when a order id is empty.
	ErrEmptyOrderID = errors.New("order id can't be empty")
	//ErrOverwriteOrderID is returned when a valid order id wants changes.
	ErrOverwriteOrderID = errors.New("customer identifier cannot overwrite")
)

func NewOrder(customerID uuid.UUID) (*Order, error) {
	if uuidIsNotEmpty(customerID) {
		return &Order{customerID: customerID}, nil
	}
	return &Order{}, ErrEmptyCustomerID
}

func (order *Order) GetID() uuid.UUID {
	return order.id
}
func (order *Order) SetID(id uuid.UUID) error {

	if uuidIsNotEmpty(order.id) {
		return ErrOverwriteOrderID
	}
	if uuidIsNotEmpty(id) {
		order.id = id
		return nil
	}

	return ErrEmptyOrderID
}
func (order *Order) GetOrderedProducts() []*OrderedProduct {
	return order.products
}
func (order *Order) SetProducts(products []*OrderedProduct) {
	order.products = products
}

func uuidIsNotEmpty(customerID uuid.UUID) bool {
	return customerID != uuid.Nil
}
