package usecase

import "github.com/KakinokiKanta/Mybrary-backend/domain"

type mockUserRepoStore struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoStore) Create(_ domain.User) (domain.User, error) {
	return *m.result, m.err
}
