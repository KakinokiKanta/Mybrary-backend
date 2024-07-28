package repositorymock

import "github.com/KakinokiKanta/Mybrary-backend/domain"

type MockUserRepoStore struct {
	domain.UserRepository

	Result *domain.User
	Err    error
}

func (m MockUserRepoStore) Create(_ domain.User) (domain.User, error) {
	return *m.Result, m.Err
}
