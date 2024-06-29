package usecases

import (
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateUserUsecase(t *testing.T) {
	successMockResult, successMockErr := domain.NewUser("Test user")
	failMockResult, failMockErr := domain.NewUser("")

	tests := []struct {
		testName string
		repository domain.UserRepository
		args CreateUserInputDTO
		expected *CreateUserOutputDTO
		expectedErr bool
	}{
		{
			testName: "Successfully create user",
			repository: mockUserRepoStore{
				result: successMockResult,
				err: successMockErr,
			},
			args: CreateUserInputDTO{
				Name: "Test user",
			},
			expected: &CreateUserOutputDTO{
				Name: "Test user",
			},
			expectedErr: false,
		},
		{
			testName: "Unsuccessfully create user",
			repository: mockUserRepoStore{
				result: failMockResult,
				err: failMockErr,
			},
			args: CreateUserInputDTO{
				Name: "",
			},
			expected: nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			uc := NewCreateUserUsecase(tt.repository)

			result, err := uc.Execute(tt.args)
			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.testName, err, tt.expectedErr)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmpopts.IgnoreFields(CreateUserOutputDTO{}, "Id", "CreatedAt"),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}
