package usecase

import "clean-arch-hicoll/pkg/domain"

type StudentUsecase struct {
	studentRepository domain.StudentRepository
}

func NewStudentUsecase(sr domain.StudentUsecase) domain.StudentUsecase {
	return &StudentUsecase{
		studentRepository: sr,
	}
}

func (su *StudentUsecase) AddNewStudent(student domain.Student) error {
	err := su.studentRepository.AddNewStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (su *StudentUsecase) GetAllStudents() ([]domain.Student, error) {
	students, err := su.studentRepository.GetAllStudents()
	if err != nil {
		return students, err
	}
	return students, err
}

func (su *StudentUsecase) GetStudentById(studentId int) (domain.Student, error) {
	student, err := su.studentRepository.GetStudentById(studentId)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (su *StudentUsecase) UpdateStudentById(student domain.Student) error {
	_, err := su.studentRepository.GetStudentById(student.Id)
	if err != nil {
		return err
	}

	err = su.studentRepository.UpdateStudentById(student)
	if err != nil {
		return err
	}
	return nil
}

func (su *StudentUsecase) DeleteStudentById(studentId int) error {
	_, err := su.studentRepository.GetStudentById(studentId)
	if err != nil {
		return err
	}

	err = su.studentRepository.DeleteStudentById(studentId)
	if err != nil {
		return err
	}
	return nil
}
