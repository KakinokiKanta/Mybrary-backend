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
		Id: createdUser,
	}
}
