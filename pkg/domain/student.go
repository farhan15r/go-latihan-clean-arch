package domain

import "clean-arch-hicoll/pkg/dto"

type Student struct {
	Id         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

type StudentRepository interface {
	AddNewStudent(student Student) error
	GetAllStudents() ([]Student, error)
	GetStudentById(studentId int) (Student, error)
	UpdateStudentById(student Student) error
	DeleteStudentById(studentId int) error
}

type StudentUsecase interface {
	AddNewStudent(student dto.StudentDTO) error
	GetAllStudents() ([]Student, error)
	GetStudentById(studentId int) (Student, error)
	UpdateStudentById(student dto.StudentDTO, studentId int) error
	DeleteStudentById(studentId int) error
}
