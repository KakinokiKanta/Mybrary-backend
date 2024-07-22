package router

import (
	"database/sql"
	"log"

	"github.com/KakinokiKanta/Mybrary-backend/api/controller"
	"github.com/KakinokiKanta/Mybrary-backend/api/middleware"
	"github.com/KakinokiKanta/Mybrary-backend/infrastracture/repository"
	"github.com/KakinokiKanta/Mybrary-backend/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	createUserUsecase := usecase.NewCreateUserUsecase(userRepository)
	loginUsecase := usecase.NewLoginUsecase(userRepository)
	createUserController := controller.NewCreateUserController(*createUserUsecase)
	jwtMiddleware, err := middleware.NewJwtMiddleware(*loginUsecase)
	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.Use(cors.Default())

	authRouter := r.Group("/auth")
	authRouter.POST("/register", createUserController.Execute)
	authRouter.POST("login", jwtMiddleware.LoginHandler)

	checkRouter := r.Group("/ping", jwtMiddleware.MiddlewareFunc())
	checkRouter.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // dockerでポート8080を指定しているため、ここでは指定しない
}
