package customer

import (
	"testing"
)

func Test_NewCustomerValidations(t *testing.T) {
	type testCase struct {
		test        string
		firstName   string
		lastName    string
		phoneNumber string
		email       string
		expectedErr error
	}

	testcase := []testCase{
		{
			test:        "empty first name validation",
			firstName:   "",
			lastName:    "Parsa",
			phoneNumber: "00989130002243",
			email:       "Parsa@gmail.com",
			expectedErr: ErrInvalidFirstName,
		},
		{
			test:        "empty last name validation",
			firstName:   "Mohamad",
			lastName:    "",
			phoneNumber: "00989130002243",
			email:       "Parsa@gmail.com",
			expectedErr: ErrInvalidLastName,
		},
		{
			test:        "empty phone number validation",
			firstName:   "Mohamad",
			lastName:    "Parsa",
			phoneNumber: "",
			email:       "Parsa@gmail.com",
			expectedErr: ErrInvalidPhoneNumber,
		},
		{
			test:        "empty email validation",
			firstName:   "Mohamad",
			lastName:    "Parsa",
			phoneNumber: "00989130002243",
			email:       "",
			expectedErr: ErrInvalidEmail,
		},
		{
			test:        "valid parameters",
			firstName:   "Mohamad",
			lastName:    "Parsa",
			phoneNumber: "00989130002243",
			email:       "Parsa@gmail.com",
			expectedErr: nil,
		},
	}

	for _, item := range testcase {
		t.Run(item.test, func(t *testing.T) {
			_, err := NewCustomer(item.firstName, item.lastName, item.phoneNumber, item.email)
			if err != item.expectedErr {
				t.Errorf("expected error :%v got %g ", item.expectedErr, err)
			}
		})
	}

}
