package cmd

import (
	"clean-arch-hicoll/config"
	"clean-arch-hicoll/pkg/controller"
	"clean-arch-hicoll/pkg/router"
	"clean-arch-hicoll/shared/db"

	"github.com/labstack/echo/v4"
)

func RunServer() {
	e := echo.New()
	g := e.Group("")

	Apply(e, g)

	e.Logger.Fatal(e.Start(":8000"))
}

func Apply(e *echo.Echo, g *echo.Group) {
	conf := config.NewConfiguration()
	db := db.NewInstanceDb(conf)
	middleware := controller.NewMiddleware(conf)

	e.Use(middleware.ErrorHandler)
	router.NewStudentRouter(e, g, db)
	router.NewUserRouter(e, g, db, middleware)
	router.NewAuthenticationRouter(e, g, db, conf)

	g.RouteNotFound("/*", middleware.NotFoundHandler)
}
