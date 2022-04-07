package order

//TODO: implement delete action.
import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrOrderNotFound is returned when a order is not found.
	ErrOrderNotFound = errors.New("the order was not found in the repository")
	// ErrFailedToAddOrder is returned when the order could not be added to the repository.
	ErrFailedToAddOrder = errors.New("failed to add the order to the repository")
	// ErrUpdateOrder is returned when the order could not be updated in the repository.
	ErrUpdateOrder = errors.New("failed to update the order in the repository")
)

//OrderRepository interface used to replace any repository to store order data.
type OrderRepository interface {
	Get(uuid.UUID) (*Order, error)
	Add(Order) error
	Update(Order) error
}
