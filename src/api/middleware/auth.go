package middleware

import (
	"time"

	"github.com/KakinokiKanta/Mybrary-backend/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func NewJwtMiddleware(uc usecase.LoginUsecase) (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "mybrary authentication",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24,
		MaxRefresh: time.Hour * 24 * 7,
		SendCookie: false,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}
		},
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var input usecase.LoginInputDTO

			if err := ctx.ShouldBind(&input); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userID, isValid := uc.Execute(input)
			if !isValid {
				return "", jwt.ErrFailedAuthentication
			}

			return userID, nil
		},
	})

	if err != nil {
		return nil, err
	}

	err = jwtMiddleware.MiddlewareInit()

	if err != nil {
		return nil, err
	}

	return jwtMiddleware, nil
}

func GetUserID(ctx *gin.Context) string {
	claims := jwt.ExtractClaims(ctx)
	userID := claims[jwt.IdentityKey]
	return userID.(string)
}
