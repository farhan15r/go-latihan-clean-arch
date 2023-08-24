package repository

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/exception"
	"database/sql"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) domain.StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

func (sr *StudentRepository) AddNewStudent(student domain.Student) error {
	query := "INSERT INTO students (fullname, address, birthdate, class, batch, school_name ) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := sr.DB.Exec(query, student.Fullname, student.Address, student.Birthdate, student.Class, student.Batch, student.SchoolName)
	if err != nil {
		return err
	}
	return nil
}

func (sr *StudentRepository) GetAllStudents() ([]domain.Student, error) {
	query := "SELECT id, fullname, address, birthdate, class, batch, school_name FROM students"

	students := []domain.Student{}

	rows, err := sr.DB.Query(query)
	if err != nil {
		return students, err
	}

	for rows.Next() {
		student := domain.Student{}

		rows.Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)

		students = append(students, student)
	}

	return students, nil
}

func (sr *StudentRepository) GetStudentById(studentId int) (domain.Student, error) {
	query := "SELECT id, fullname, address, birthdate, class, batch, school_name FROM students WHERE id = $1 LIMIT 1"

	student := domain.Student{}

	rows, err := sr.DB.Query(query, studentId)
	if err != nil {
		return student, err
	}

	if rows.Next() {
		rows.Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
	} else {
		return student, exception.NewNotFoundError("student is not found")
	}

	return student, nil
}

func (sr *StudentRepository) UpdateStudentById(student domain.Student) error {
	query := `UPDATE students SET
		fullname = $1, 
		address = $2, 
		birthdate = $3, 
		class = $4, 
		batch = $5, 
		school_name	= $6
		WHERE id = $7`

	_, err := sr.DB.Exec(query, student.Fullname, student.Address, student.Birthdate, student.Class, student.Batch, student.SchoolName, student.Id)

	if err != nil {
		return err
	}
	return nil
}

func (sr *StudentRepository) DeleteStudentById(studentId int) error {
	query := "DELETE FROM students WHERE id = $1"

	_, err := sr.DB.Exec(query, studentId)

	if err != nil {
		return err
	}
	return nil
}
