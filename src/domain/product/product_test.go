package product

import (
	"testing"
)

func Test_NewCustomerValidations(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		model       string
		companyName string
		description string
		expectedErr error
	}

	testcase := []testCase{
		{
			test:        "empty commdity name validation",
			name:        "",
			model:       "SE 2021",
			companyName: "Apple",
			description: "",
			expectedErr: ErrInvalidCommodityName,
		},
		{
			test:        "empty commdity model validation",
			name:        "Apple Watch",
			model:       "",
			companyName: "Apple",
			description: "",
			expectedErr: ErrInvalidCommodityModel,
		},
		{
			test:        "empty commdity company name validation",
			name:        "Apple Watch",
			model:       "SE 2021",
			companyName: "",
			description: "",
			expectedErr: ErrInvalidCommodityCompanyName,
		},
		{
			test:        "empty description validation",
			name:        "Apple Watch",
			model:       "SE 2021",
			companyName: "Apple",
			description: "",
			expectedErr: nil,
		},
		{
			test:        "empty commdity model validation",
			name:        "Apple Watch",
			model:       "SE 2021",
			companyName: "Apple",
			description: "",
			expectedErr: nil,
		},
	}

	for _, item := range testcase {
		t.Run(item.test, func(t *testing.T) {
			_, err := NewProduct(item.name, item.model, item.companyName, item.description)
			if err != item.expectedErr {
				t.Errorf("expected error :%v got %g ", item.expectedErr, err)
			}
		})
	}

}
