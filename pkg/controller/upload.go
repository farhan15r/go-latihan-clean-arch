package controller

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/response"

	"github.com/labstack/echo/v4"
)

type UploadController struct {
	Uu domain.UploadUsecase
}

func (uc *UploadController) UploadNewFile(c echo.Context) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return err
	}

	file, err := formFile.Open()
	fileName := formFile.Filename
	if err != nil {
		return err
	}

	err = uc.Uu.UploadNewFile(file, fileName)
	if err != nil {
		return err
	}

	return response.SetResponse(c, 201, "success", nil)
}
