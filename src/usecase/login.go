package usecase

import (
	"github.com/KakinokiKanta/Mybrary-backend/domain"
)

type LoginUsecase struct {
	userRepo domain.UserRepository
}

type LoginInputDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=64"`
}

func NewLoginUsecase(userRepo domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		userRepo: userRepo,
	}
}

/*
 * ここではトークンの生成は行わず、ログインの成否だけを判定する
 * これにより、トークン生成を行う外部パッケージに依存しない
 * また、認証手段を変更しやすくしている
 */
func (uc LoginUsecase) Execute(input LoginInputDTO) (bool, error) {
	// TODO: DB内に同一メールアドレスが複数存在している場合のために、以下の処理をループさせる？
	// メールアドレスでDB内を検索しユーザ情報を取得
	user, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	return input.Password == user.Password(), nil
}
