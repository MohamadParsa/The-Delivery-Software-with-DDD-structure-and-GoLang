package memory

import (
	"domain/order"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetOrder(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	fakeOrder := getOrder(t)
	id := fakeOrder.GetID()
	repository := MemoryRepository{orders: map[uuid.UUID]order.Order{}}
	repository.orders[id] = *fakeOrder

	testCases := []testCase{
		{
			name:        "find order with invalid id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: order.ErrOrderNotFound,
		},
		{
			name:        "find order with valid id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			_, errOfGet := repository.Get(testcase.id)
			if errOfGet != testcase.expectedErr {
				t.Errorf("expected error :%v , got %v", testcase.expectedErr, errOfGet)
			}
		})
	}
}

func TestMemory_AddOrder(t *testing.T) {
	type testCase struct {
		name        string
		newOrder    order.Order
		expectedErr error
	}
	validOrder := getOrder(t)
	testCases := []testCase{
		{
			name:        "add valid order",
			newOrder:    *validOrder,
			expectedErr: nil,
		},
		{
			name:        "add duplicated order",
			newOrder:    *validOrder,
			expectedErr: order.ErrFailedToAddOrder,
		},
	}
	memoryRepository := New()
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Add(testcase.newOrder)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error :%v, got %v", testcase.expectedErr, errOfAdd)
			}

			id := testcase.newOrder.GetID()

			found, errGet := memoryRepository.Get(id)
			checkError(errGet, t)
			if found.GetID() != id {
				t.Errorf("expected %v, got %v", id, found.GetID())
			}
		})
	}
}

func TestMemory_UpdateOrder(t *testing.T) {
	type testCase struct {
		name        string
		newOrder    order.Order
		expectedErr error
	}

	validOrder := getOrder(t)
	invalidOrder := getOrder(t)
	memoryRepository := New()
	errOfAdd := memoryRepository.Add(*validOrder)
	checkError(errOfAdd, t)

	testCases := []testCase{
		{
			name:        "update valid order",
			newOrder:    *validOrder,
			expectedErr: nil,
		},
		{
			name:        "update invalid order",
			newOrder:    *invalidOrder,
			expectedErr: order.ErrUpdateOrder,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Update(testcase.newOrder)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error :%v, got %v", testcase.expectedErr, errOfAdd)
			}

			errUpdate := memoryRepository.Update(testcase.newOrder)

			if errUpdate != testcase.expectedErr {
				t.Errorf("expected %v, got %v", testcase.expectedErr, errUpdate)
			}
		})
	}
}

func getOrder(t *testing.T) *order.Order {
	customerID, err := uuid.NewUUID()
	checkError(err, t)
	validOrder, err := order.NewOrder(customerID)
	checkError(err, t)
	orderID, err := uuid.NewUUID()
	checkError(err, t)
	err = validOrder.SetID(orderID)
	checkError(err, t)
	return validOrder
}
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
