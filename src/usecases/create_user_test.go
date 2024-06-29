package usecases

import (
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
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
}
