package memory

import (
	"domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	fakeCustomer := getCustomer(t)
	id := fakeCustomer.GetID()
	repository := MemoryRepository{customers: map[uuid.UUID]customer.Customer{}}
	repository.customers[id] = *fakeCustomer

	testCases := []testCase{
		{
			name:        "find customer with invalid id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "find customer with valid id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			_, errOfGet := repository.Get(testcase.id)
			if errOfGet != testcase.expectedErr {
				t.Errorf("expected error %v , got %v", testcase.expectedErr, errOfGet)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		newCustomer customer.Customer
		expectedErr error
	}
	validCustomer := getCustomer(t)
	testCases := []testCase{
		{
			name:        "add valid customer",
			newCustomer: *validCustomer,
			expectedErr: nil,
		},
		{
			name:        "add duplicated customer",
			newCustomer: *validCustomer,
			expectedErr: customer.ErrFailedToAddCustomer,
		},
	}
	memoryRepository := New()
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Add(testcase.newCustomer)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error %v, got %v", testcase.expectedErr, errOfAdd)
			}

			id := testcase.newCustomer.GetID()

			found, errGet := memoryRepository.Get(id)
			checkError(errGet, t)
			if found.GetID() != id {
				t.Errorf("expected %v, got %v", id, found.GetID())
			}
		})
	}
}

func TestMemory_UpdateCustomer(t *testing.T) {
	type testCase struct {
		name        string
		newCustomer customer.Customer
		expectedErr error
	}

	validCustomer := getCustomer(t)
	invalidCustomer := getCustomer(t)
	memoryRepository := New()
	errOfAdd := memoryRepository.Add(*validCustomer)
	checkError(errOfAdd, t)

	testCases := []testCase{
		{
			name:        "update valid customer",
			newCustomer: *validCustomer,
			expectedErr: nil,
		},
		{
			name:        "update invalid customer",
			newCustomer: *invalidCustomer,
			expectedErr: customer.ErrUpdateCustomer,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {

			errOfAdd := memoryRepository.Update(testcase.newCustomer)
			if errOfAdd != testcase.expectedErr {
				t.Errorf("expected error %v, got %v", testcase.expectedErr, errOfAdd)
			}

			errUpdate := memoryRepository.Update(testcase.newCustomer)

			if errUpdate != testcase.expectedErr {
				t.Errorf("expected %v, got %v", testcase.expectedErr, errUpdate)
			}
		})
	}
}

func getCustomer(t *testing.T) *customer.Customer {
	validCustomer, err := customer.NewCustomer("Mohamad", "Parsa", "00989122212221", "Parsa@gmail.com")
	checkError(err, t)
	id, err := uuid.NewUUID()
	checkError(err, t)
	err = validCustomer.SetID(id)
	checkError(err, t)
	return validCustomer
}
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
