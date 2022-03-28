//Package memory is a simple implementation of in-memory storage.
package memory

import (
	"domain/product"
	"sync"

	"github.com/google/uuid"
)
//MemoryRepository is data repository.
type MemoryRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}
//New function return an instance of MemoryRepository.
func New() *MemoryRepository {
	return &MemoryRepository{products: make(map[uuid.UUID]product.Product, 0)}
}
//GetAll function returns all Product object.
func (memoryRepository *MemoryRepository) GetAll() (map[uuid.UUID]product.Product, error) {
	return memoryRepository.products, nil
}
//Get function search for a product by id and return a Product object.
func (memoryRepository *MemoryRepository) Get(id uuid.UUID) (product.Product, error) {
	if product, ok := memoryRepository.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}
//Add function add a new product to repository.
func (memoryRepository *MemoryRepository) Add(newProduct product.Product) error {
	if memoryRepository.products == nil {
		memoryRepository.Mutex.Lock()
		memoryRepository.products = make(map[uuid.UUID]product.Product)
		memoryRepository.Mutex.Unlock()
	}
	if memoryRepository.productExists(newProduct){
		return product.ErrFailedToAddProduct
	}

	memoryRepository.Mutex.Lock()
	memoryRepository.products[newProduct.GetID()] = newProduct
	memoryRepository.Mutex.Unlock()
	return nil
}
//Update function update a new product to repository.
func (memoryRepository *MemoryRepository) Update(newProduct product.Product) error {
	if memoryRepository.products == nil {
		memoryRepository.Mutex.Lock()
		memoryRepository.products = make(map[uuid.UUID]product.Product)
		memoryRepository.Mutex.Unlock()
	}
	if !memoryRepository.productExists(newProduct) {
		return product.ErrUpdateProduct
	}

	memoryRepository.Mutex.Lock()
	memoryRepository.products[newProduct.GetID()] = newProduct
	memoryRepository.Mutex.Unlock()
	return nil
}
func (memory *MemoryRepository) productExists(product product.Product) bool {
	_, exists := memory.products[product.GetID()]
	return exists
}
