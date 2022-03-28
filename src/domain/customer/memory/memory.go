//Package memory is a simple implementation of in-memory storage.
package memory

import (
	"domain/customer"
	"sync"

	"github.com/google/uuid"
)

//MemoryRepository is data repository.
type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

//New function return an instance of MemoryRepository.
func New() *MemoryRepository {
	return &MemoryRepository{customers: make(map[uuid.UUID]customer.Customer, 0)}
}

//Get function search for a customer by id and return a Customer object.
func (memory *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := memory.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

//Add function add a new customer to repository.
func (memory *MemoryRepository) Add(newcustomer customer.Customer) error {
	if memory.customers == nil {
		memory.Lock()
		memory.customers = make(map[uuid.UUID]customer.Customer)
		memory.Unlock()
	}
	if memory.customerExists(newcustomer) {
		return customer.ErrFailedToAddCustomer
	}
	memory.Lock()
	memory.customers[newcustomer.GetID()] = newcustomer
	memory.Unlock()
	return nil
}

//Update function update a new customer to repository.
func (memory *MemoryRepository) Update(newcustomer customer.Customer) error {
	if memory.customers == nil {
		memory.Lock()
		memory.customers = make(map[uuid.UUID]customer.Customer)
		memory.Unlock()
	}
	if !memory.customerExists(newcustomer) {
		return customer.ErrUpdateCustomer
	}
	memory.Lock()
	memory.customers[newcustomer.GetID()] = newcustomer
	memory.Unlock()
	return nil
}

func (memory *MemoryRepository) customerExists(customer customer.Customer) bool {
	_, exists := memory.customers[customer.GetID()]
	return exists
}
