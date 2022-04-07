package order

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewOrder(t *testing.T) {

	type TestCase struct {
		test        string
		customerID  uuid.UUID
		expectedErr error
	}
	customerID, err := uuid.NewUUID()
	checkError(err, t)
	testCases := []TestCase{
		{
			test:        "empty customer id",
			expectedErr: ErrEmptyCustomerID,
		},
		{
			test:        "new order with valid data",
			customerID:  customerID,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := NewOrder(testCase.customerID)
			if err != testCase.expectedErr {
				t.Errorf("expected error :%v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
