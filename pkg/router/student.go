package router

import (
	"clean-arch-hicoll/pkg/controller"
	"clean-arch-hicoll/pkg/repository"
	"clean-arch-hicoll/pkg/usecase"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewStudentRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	sr := repository.NewStudentRepository(db)
	su := usecase.NewStudentUsecase(sr)
	sc := &controller.StudentControler{
		StudentUsecase: su,
	}

	e.POST("/students", sc.PostStudent)
	e.GET("/students", sc.GetStudents)
	e.GET("/students/:studentId", sc.GetStudent)
	e.PUT("/students/:studentId", sc.PutStudent)
	e.DELETE("/students/:studentId", sc.DeleteStudent)
}
