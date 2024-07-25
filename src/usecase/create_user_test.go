package usecase

import (
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	repositorymock "github.com/KakinokiKanta/Mybrary-backend/usecase/repository_mock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateUserUsecase(t *testing.T) {
	successMockResult, successMockErr := domain.NewUser("Test user", "test@gmail.com", "abcdefg1234AABBCCDD")
	failMockResult, failMockErr := domain.NewUser("", "", "")

	tests := []struct {
		testName string
		repository domain.UserRepository
		args CreateUserInputDTO
		expected *CreateUserOutputDTO
		expectedErr bool
	}{
		{
			testName: "Successfully create user",
			repository: repositorymock.MockUserRepoStore{
				Result: successMockResult,
				Err: successMockErr,
			},
			args: CreateUserInputDTO{
				Name: "Test user",
				Email: "test@gmail.com",
				Password: "abcdefg1234AABBCCDD",
			},
			expected: &CreateUserOutputDTO{
				Name: "Test user",
				Email: "test@gmail.com",
				Password: "abcdefg1234AABBCCDD",
			},
			expectedErr: false,
		},
		{
			testName: "Unsuccessfully create user",
			repository: repositorymock.MockUserRepoStore{
				Result: failMockResult,
				Err: failMockErr,
			},
			args: CreateUserInputDTO{
				Name: "",
				Email: "",
				Password: "",
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
