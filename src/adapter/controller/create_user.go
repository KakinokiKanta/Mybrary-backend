package controller

import (
	"github.com/KakinokiKanta/Mybrary-backend/usecase"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	uc usecase.CreateUserUsecase
}

func NewCreateUserController(uc usecase.CreateUserUsecase) CreateUserController {
	return CreateUserController{
		uc: uc,
	}
}

func (con CreateUserController) Execute(ctx *gin.Context) {
	var input usecase.CreateUserInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(500, gin.H{"error": "error"})
		return
	}
}
