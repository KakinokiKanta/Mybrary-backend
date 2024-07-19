package usecase

import (
	"database/sql"
	"errors"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type FindUserUsecase struct {
	userRepo domain.UserRepository
}

type FindUserInputDTO struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=64"`
}

type FindUserOutputDTO struct {
	Id       domain.UserID `json:"id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"passdomain`
}

func NewFindUserUsecase(userRepo domain.UserRepository) *FindUserUsecase {
	return &FindUserUsecase{
		userRepo: userRepo,
	}
}

// ログイン機能
func (uc FindUserUsecase) Execute(input FindUserInputDTO) (*FindUserOutputDTO, error) {
	// Userドメインを生成
	user, err := domain.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	// 同一メールアドレスがDB内に存在しないかチェック
	_, err = uc.userRepo.FindByEmail(user.Email())
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
