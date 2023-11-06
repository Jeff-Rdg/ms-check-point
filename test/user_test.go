package test

import (
	"errors"
	"ms-control-point/internal/entity"
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	type TestCase struct {
		testCaseName string
		name         string
		email        string
		password     string
		expectedErr  error
	}

	testCases := []TestCase{
		{
			testCaseName: "Valid user",
			name:         "John Doe",
			email:        "jonh@gmail.com",
			password:     "123456",
			expectedErr:  nil,
		},
		{
			testCaseName: "required Name field",
			name:         "",
			email:        "jonh@gmail.com",
			password:     "123456",
			expectedErr:  entity.ErrNameIsRequired,
		},
		{
			testCaseName: "required email field",
			name:         "John Doe",
			email:        "",
			password:     "123456",
			expectedErr:  entity.ErrEmailIsRequired,
		},
		{
			testCaseName: "email invalid",
			name:         "John Doe",
			email:        "jota!gmail.com",
			password:     "123456",
			expectedErr:  entity.ErrInvalidEmail,
		},
		{
			testCaseName: "required password field",
			name:         "John Doe",
			email:        "jonh@gmail.com",
			password:     "",
			expectedErr:  entity.ErrPasswordIsRequired,
		},
	}

	for _, test := range testCases {
		t.Run(test.testCaseName, func(t *testing.T) {
			_, err := entity.NewUser(test.name, test.email, test.password)
			if !errors.Is(err, test.expectedErr) {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	type TestCase struct {
		testCaseName    string
		user            *entity.User
		comparePassword string
		expectedResult  bool
	}

	testCases := []TestCase{
		{
			testCaseName:    "Valid password",
			user:            user,
			comparePassword: "123456",
			expectedResult:  true,
		},
		{
			testCaseName:    "Invalid password",
			user:            user,
			comparePassword: "1234567",
			expectedResult:  false,
		},
	}

	for _, test := range testCases {
		t.Run(test.testCaseName, func(t *testing.T) {
			validatePassword := test.user.ValidatePassword(test.comparePassword)
			if !validatePassword == test.expectedResult {
				t.Errorf("expected error %v, got %v", test.expectedResult, validatePassword)
			}
		})
	}

}
