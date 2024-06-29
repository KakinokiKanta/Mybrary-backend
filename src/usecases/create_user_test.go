package usecases

import (
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockUserRepoStore struct {
	domain.UserRepository

	result *domain.User
	err error
}

func (m mockUserRepoStore) Create(_ domain.User) (domain.User, error) {
	return *m.result, m.err
}

func TestCreateUserUsecase(t *testing.T) {
	repositoryResult, repositoryErr := domain.NewUser("Test user")

	tests := []struct {
		testName string
		repository domain.UserRepository
		args CreateUserInputDTO
		expected *CreateUserOutputDTO
		expectedErr bool
	}{
		{
			testName: "Create user successful",
			repository: mockUserRepoStore{
				result: repositoryResult,
				err: repositoryErr,
			},
			args: CreateUserInputDTO{
				Name: "Test user",
			},
			expected: &CreateUserOutputDTO{
				Name: "Test user",
			},
			expectedErr: false,
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
