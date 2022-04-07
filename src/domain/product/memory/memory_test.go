package memory

import (
	"domain/product"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetProduct(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	newProduct := getProduct(t)
	id := newProduct.GetID()
	repository := MemoryRepository{products: map[uuid.UUID]*product.Product{}}
	repository.products[id] = newProduct

	testCases := []testCase{
		{
			name:        "find product with invalid id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: product.ErrProductNotFound,
		},
		{
			name:        "find product with valid id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			_, errOfGet := repository.Get(testcase.id)
			if errOfGet != testcase.expectedErr {
				t.Errorf("expected error: %v , got %v", testcase.expectedErr, errOfGet)
			}
		})
	}
}

func TestMemory_AddProduct(t *testing.T) {
	type testCase struct {
		name        string
		newProduct  product.Product
		expectedErr error
	}
	validProduct := getProduct(t)
	testCases := []testCase{
		{
			name:        "add valid product",
			newProduct:  *validProduct,
			expectedErr: nil,
		},
		{
			name:        "add duplicated product",
			newProduct:  *validProduct,
			expectedErr: product.ErrFailedToAddProduct,
		},
	}
	memoryRepository := New()
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Add(testcase.newProduct)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error :%v, got %v", testcase.expectedErr, errOfAdd)
			}

			id := testcase.newProduct.GetID()

			found, errGet := memoryRepository.Get(id)
			checkError(errGet, t)
			if found.GetID() != id {
				t.Errorf("expected %v, got %v", id, found.GetID())
			}
		})
	}
}

func TestMemory_UpdateProduct(t *testing.T) {
	type testCase struct {
		name        string
		newProduct  product.Product
		expectedErr error
	}

	validProduct := getProduct(t)
	invalidProduct := getProduct(t)
	memoryRepository := New()
	errOfAdd := memoryRepository.Add(*validProduct)
	checkError(errOfAdd, t)

	testCases := []testCase{
		{
			name:        "update valid product",
			newProduct:  *validProduct,
			expectedErr: nil,
		},
		{
			name:        "update invalid product",
			newProduct:  *invalidProduct,
			expectedErr: product.ErrUpdateProduct,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Update(testcase.newProduct)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error :%v, got %v", testcase.expectedErr, errOfAdd)
			}

			errUpdate := memoryRepository.Update(testcase.newProduct)

			if errUpdate != testcase.expectedErr {
				t.Errorf("expected %v, got %v", testcase.expectedErr, errUpdate)
			}
		})
	}
}

func getProduct(t *testing.T) *product.Product {
	product, err := product.NewProduct("Apple Watch", "SE 2021", "Apple", "")
	checkError(err, t)
	id, err := uuid.NewUUID()
	checkError(err, t)
	err = product.SetID(id)
	checkError(err, t)
	return product
}
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
