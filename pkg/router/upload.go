package router

import (
	"clean-arch-hicoll/pkg/controller"
	"clean-arch-hicoll/pkg/repository"
	"clean-arch-hicoll/pkg/repository/storage"
	"clean-arch-hicoll/pkg/usecase"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func NewUploadRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	ur := repository.NewUploadRepository(db)
	us := storage.NewUploadStorage()
	uu := usecase.NewUploadUsecase(ur, us)
	uc := &controller.UploadController{
		Uu: uu,
	}

	e.POST("/uploads", uc.UploadNewFile)
}
