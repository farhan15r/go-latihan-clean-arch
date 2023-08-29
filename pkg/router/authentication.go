package router

import (
	"clean-arch-hicoll/config"
	"clean-arch-hicoll/pkg/controller"
	"clean-arch-hicoll/pkg/repository"
	"clean-arch-hicoll/pkg/usecase"
	"clean-arch-hicoll/shared/token"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewAuthenticationRouter(e *echo.Echo, g *echo.Group, db *sql.DB, conf *config.Configuration) {
	tm := token.NewTokenManager(conf)
	ar := repository.NewAuthRepository(db)
	ur := repository.NewUserRepository(db)
	au := usecase.NewAuthUsecase(tm, ar, ur)
	ac := &controller.AuthController{
		AuthUsecase: au,
	}

	e.POST("/authentications", ac.PostAuth)
	e.PUT("/authentications", ac.PutAuth)
	e.DELETE("/authentications", ac.DeleteAuth)
}
