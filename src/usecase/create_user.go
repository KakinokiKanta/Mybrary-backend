package usecase

import (
	"time"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type CreateUserUsecase struct {
	userRepo domain.UserRepository
}

type CreateUserInputDTO struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=64"`
}

type CreateUserOutputDTO struct {
	Id domain.UserID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func NewCreateUserUsecase(userRepo domain.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo: userRepo,
	}
}

func (uc CreateUserUsecase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	// Userドメインを生成
	user, err := domain.NewUser(input.Name)
	if err != nil {
		return nil, err
	}

	// Userドメインのリポジトリを用いて、Userの永続化
	createdUser, err := uc.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		Id: createdUser.ID(),
		Name: createdUser.Name(),
		CreatedAt: createdUser.CreatedAt().Format(time.RFC3339),
	}, nil
}
