package repository

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/shared/exception"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) ValidateUsernameAvailable(username string) error {
	query := "SELECT id FROM users WHERE username = $1 LIMIT 1"

	rows := ur.DB.QueryRow(query, username)

	var id int
	err := rows.Scan(&id)
	if err != nil {
		return nil
	}
	return exception.NewClientError("username is already used", "username is already used")
}

func (ur *UserRepository) ValidateEmailAvailable(email string) error {
	query := "SELECT id FROM users WHERE email = $1 LIMIT 1"

	rows := ur.DB.QueryRow(query, email)

	var id int
	err := rows.Scan(&id)
	if err != nil {
		return nil
	}
	return exception.NewClientError("email is already used", "email is already used")
}

func (ur *UserRepository) AddNewUser(user domain.User) error {
	query := "INSERT INTO users (username, email, password, address) VALUES ($1, $2, $3, $4)"

	_, err := ur.DB.Exec(query, user.Username, user.Email, user.Password, user.Address)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserById(userId int) (domain.User, error) {
	query := "SELECT id, username, email, address FROM users WHERE id = $1 LIMIT 1"

	row := ur.DB.QueryRow(query, userId)

	user := domain.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Address)
	if err != nil {
		return user, exception.NewNotFoundError("user not found", "user not found")
	}

	return user, nil
}

func (ur *UserRepository) GetAllUsers() ([]domain.User, error) {
	query := "SELECT id, username, email, address FROM users"
	users := []domain.User{}

	rows, err := ur.DB.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := domain.User{}
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.Address)

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) UpdateUserById(user domain.User) error {
	query := `UPDATE users SET
		username = $1,
		email = $2,
		password = $3,
		address = $4
		WHERE id = $5
	`

	_, err := ur.DB.Exec(query, user.Username, user.Email, user.Password, user.Address, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUserById(userId int) error {
	query := "DELETE FROM users WHERE id = $1"

	_, err := ur.DB.Exec(query, userId)

	if err != nil {
		return err
	}
	return nil
}
