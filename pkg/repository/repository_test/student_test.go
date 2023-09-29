package repository_test

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func truncateStudentTable() {
	_, err := dbConn.Exec("TRUNCATE TABLE students")
	if err != nil {
		panic(err)
	}
}

var student = domain.Student{
	Fullname:   "Joko",
	Address:    "Jakarta",
	Birthdate:  "1990-01-01T00:00:00Z",
	Class:      "10",
	SchoolName: "SMA 1",
	Batch:      2020,
}

func TestStudentRepository_Add(t *testing.T) {
	defer truncateStudentTable()
	t.Run("should success add student", func(t *testing.T) {

		repo := repository.NewStudentRepository(dbConn)
		err := repo.AddNewStudent(student)

		assert.NoError(t, err)

		row := dbConn.QueryRow("SELECT id, fullname, address, birthdate, class, batch, school_name birth FROM students WHERE fullname = $1", student.Fullname)
		var result domain.Student
		err = row.Scan(&result.Id, &result.Fullname, &result.Address, &result.Birthdate, &result.Class, &result.Batch, &result.SchoolName)

		assert.NoError(t, err)
		assert.Equal(t, student.Fullname, result.Fullname)
		assert.Equal(t, student.Address, result.Address)
		assert.Equal(t, student.Birthdate, result.Birthdate)
		assert.Equal(t, student.Class, result.Class)
		assert.Equal(t, student.Batch, result.Batch)
		assert.Equal(t, student.SchoolName, result.SchoolName)
	})
}

func TestStudentRepository_GetAll(t *testing.T) {
	defer truncateStudentTable()
	t.Run("should success get all students length 1", func(t *testing.T) {
		repo := repository.NewStudentRepository(dbConn)

		repo.AddNewStudent(student)

		result, err := repo.GetAllStudents()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("should success get all students length 2", func(t *testing.T) {
		repo := repository.NewStudentRepository(dbConn)

		repo.AddNewStudent(student)

		result, err := repo.GetAllStudents()

		assert.NoError(t, err)
		assert.Equal(t, 2, len(result))
	})
}

func TestStudentRepository_GetById(t *testing.T) {
	defer truncateStudentTable()
	t.Run("should success get student by id", func(t *testing.T) {
		repo := repository.NewStudentRepository(dbConn)

		repo.AddNewStudent(student)

		studentId := 0
		dbConn.QueryRow("SELECT id FROM students WHERE fullname = $1 LIMIT 1", student.Fullname).Scan(&studentId)

		result, err := repo.GetStudentById(studentId)
		assert.NoError(t, err)
		assert.Equal(t, student.Fullname, result.Fullname)
		assert.Equal(t, student.Address, result.Address)
		assert.Equal(t, student.Birthdate, result.Birthdate)
		assert.Equal(t, student.Class, result.Class)
		assert.Equal(t, student.Batch, result.Batch)
		assert.Equal(t, student.SchoolName, result.SchoolName)
	})
}

func TestStudentRepository_Update(t *testing.T) {
	defer truncateStudentTable()

	t.Run("should success update student", func(t *testing.T) {
		repo := repository.NewStudentRepository(dbConn)

		repo.AddNewStudent(student)

		dbConn.QueryRow("SELECT id FROM students WHERE fullname = $1 LIMIT 1", student.Fullname).Scan(&student.Id)

		student.Fullname = "Joko Susilo" // update fullname
		err := repo.UpdateStudentById(student)

		assert.NoError(t, err)

		resultStudent, _ := repo.GetStudentById(student.Id)

		assert.Equal(t, student.Fullname, resultStudent.Fullname)
		assert.Equal(t, student.Address, resultStudent.Address)
		assert.Equal(t, student.Birthdate, resultStudent.Birthdate)
		assert.Equal(t, student.Class, resultStudent.Class)
		assert.Equal(t, student.Batch, resultStudent.Batch)
		assert.Equal(t, student.SchoolName, resultStudent.SchoolName)
	})
}

func TestStudentRepository_Delete(t *testing.T) {
	defer truncateStudentTable()

	t.Run("should success delete student", func(t *testing.T) {
		repo := repository.NewStudentRepository(dbConn)

		repo.AddNewStudent(student)

		dbConn.QueryRow("SELECT id FROM students WHERE fullname = $1 LIMIT 1", student.Fullname).Scan(&student.Id)

		err := repo.DeleteStudentById(student.Id)

		assert.NoError(t, err)

		_, err = repo.GetStudentById(student.Id)

		assert.Error(t, err)
	})
}
