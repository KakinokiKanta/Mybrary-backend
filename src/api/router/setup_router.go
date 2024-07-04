package router

import (
	"database/sql"

	"github.com/KakinokiKanta/Mybrary-backend/api/controller"
	"github.com/KakinokiKanta/Mybrary-backend/infrastracture/repository"
	"github.com/KakinokiKanta/Mybrary-backend/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(db *sql.DB) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository)
	createUserController := controller.NewCreateUserController(*createUserUsecase)

	r := gin.Default()
	r.Use(cors.Default())

	userRouter := r.Group("/user")
	userRouter.POST("", createUserController.Execute)

	return r
}
