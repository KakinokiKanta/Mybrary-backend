package controller

import "github.com/KakinokiKanta/Mybrary-backend/usecase"

type CreateUserController struct {
	uc usecase.CreateUserUsecase
}

func NewCreateUserController(uc usecase.CreateUserUsecase) CreateUserController {
	return CreateUserController{
		uc: uc,
	}
}
