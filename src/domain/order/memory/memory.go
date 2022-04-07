package memory

import (
	"domain/order"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	orders map[uuid.UUID]order.Order
	sync.Mutex
}

//New function return an instance of MemoryRepository.
func New() *MemoryRepository {
	return &MemoryRepository{orders: make(map[uuid.UUID]order.Order)}
}

//Get function search for a order by id and return a Order object.
func (memoryRepository *MemoryRepository) Get(id uuid.UUID) (*order.Order, error) {
	if order, ok := memoryRepository.orders[id]; ok {
		return &order, nil
	}
	return &order.Order{}, order.ErrOrderNotFound
}

//Add function add a new order to repository.
func (memoryRepository *MemoryRepository) Add(newOrder order.Order) error {
	memoryRepository.createOrderMapIfNotExists()

	if memoryRepository.orderExists(newOrder) || uuidIsEmpty(newOrder.GetID()) {
		return order.ErrFailedToAddOrder
	}
	memoryRepository.setOrderIntoMap(newOrder)
	return nil
}

//Update function update a new order to repository.
func (memoryRepository *MemoryRepository) Update(newOrder order.Order) error {
	memoryRepository.createOrderMapIfNotExists()

	if !memoryRepository.orderExists(newOrder) || uuidIsEmpty(newOrder.GetID()) {
		return order.ErrUpdateOrder
	}
	memoryRepository.setOrderIntoMap(newOrder)
	return nil
}
func uuidIsEmpty(orderID uuid.UUID) bool {
	return orderID == uuid.Nil
}
func (memoryRepository *MemoryRepository) orderExists(order order.Order) bool {
	_, exists := memoryRepository.orders[order.GetID()]
	return exists
}
func (memoryRepository *MemoryRepository) createOrderMapIfNotExists() {
	if memoryRepository.orders == nil {
		memoryRepository.Lock()
		memoryRepository.orders = make(map[uuid.UUID]order.Order)
		memoryRepository.Unlock()
	}
}
func (memoryRepository *MemoryRepository) setOrderIntoMap(neworder order.Order) {
	memoryRepository.Lock()
	memoryRepository.orders[neworder.GetID()] = neworder
	memoryRepository.Unlock()
}
