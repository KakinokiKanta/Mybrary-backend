package controller

import (
	"net/http"

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
	// JSONリクエストを受け取る
	var input usecase.CreateUserInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 新規ユーザ登録のユースケースを実行
	registeredUser, err := con.uc.Execute(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": registeredUser})
}
