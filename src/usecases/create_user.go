package usecases

import "github.com/KakinokiKanta/Mybrary-backend/domain"

type CreateUserUsecase struct {
	userRepo domain.UserRepository
}

type CreateUserInputDTO struct {
	Name string `json:"name"`
}

type CreateUserOutputDTO struct {
	Id domain.UserID `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func NewCreateUserUsecase(userRepo domain.UserRepository) CreateUserUsecase {
	return CreateUserUsecase{
		userRepo: userRepo,
	}
}
