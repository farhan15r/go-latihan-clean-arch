package router

import (
	"clean-arch-hicoll/pkg/controller"
	"clean-arch-hicoll/pkg/repository"
	"clean-arch-hicoll/pkg/usecase"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo, g *echo.Group, db *sql.DB, middleware *controller.Middleware) {
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := &controller.UserController{
		UserUsecase: uu,
	}

	e.POST("/users", uc.PostUser)
	e.GET("/users", uc.GetUsers)
	e.GET("/users/:userId", uc.GetUser, middleware.RequireLogin)
	e.PUT("/users/:userId", uc.PutUser)
	e.DELETE("/users/:userId", uc.DeleteUser)
}
