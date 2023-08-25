package controller

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/dto"
	"clean-arch-hicoll/shared/exception"
	"clean-arch-hicoll/shared/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StudentControler struct {
	StudentUsecase domain.StudentUsecase
}

func (sc *StudentControler) PostStudent(c echo.Context) error {
	student := dto.StudentDTO{}

	err := c.Bind(&student)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	err = sc.StudentUsecase.AddNewStudent(student)

	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusCreated, "success", nil)
}

func (sc *StudentControler) GetStudents(c echo.Context) error {
	students, err := sc.StudentUsecase.GetAllStudents()
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", students)
}

func (sc *StudentControler) GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	student, err := sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", student)
}

func (sc *StudentControler) PutStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	student := dto.StudentDTO{}

	err = c.Bind(&student)
	if err != nil {
		return exception.NewClientError("payload is not valid", err)
	}

	err = sc.StudentUsecase.UpdateStudentById(student, id)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", nil)
}

func (sc *StudentControler) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		return exception.NewClientError("param is not valid", "param is not valid")
	}

	err = sc.StudentUsecase.DeleteStudentById(id)
	if err != nil {
		return err
	}

	return response.SetResponse(c, http.StatusOK, "success", nil)
}
