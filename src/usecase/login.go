package usecase

import (
	"database/sql"
	"errors"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type LoginUsecase struct {
	userRepo domain.UserRepository
}

type LoginInputDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=64"`
}

type LoginOutputDTO struct {
	Token string `json:"token"`
}

func NewLoginUsecase(userRepo domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		userRepo: userRepo,
	}
}

func (uc LoginUsecase) Execute(input LoginInputDTO) (*LoginOutputDTO, error) {
	// 同一メールアドレスがDB内に存在しないかチェック
	_, err = uc.userRepo.FindByEmail(input.Email)
	if err == nil {
		return nil, errors.New("this email address is already registered")
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	// Userドメインのリポジトリを用いて、Userの永続化
	createdUser, err := uc.userRepo.FindByEmail(/*string*/)
	if err != nil {
		return nil, err
	}

	return &FindUserOutputDTO{
		Id:       createdUser.ID(),
		Name:     createdUser.Name(),
		Email:    createdUser.Email(),
		Password: createdUser.Password(),
	}, nil
}
