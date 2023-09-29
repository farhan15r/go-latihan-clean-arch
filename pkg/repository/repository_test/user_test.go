package repository_test

import (
	"clean-arch-hicoll/pkg/domain"
	"clean-arch-hicoll/pkg/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func truncateUsersTable() {
	_, err := dbConn.Exec("TRUNCATE TABLE users")
	if err != nil {
		panic(err)
	}
}

var user = domain.User{
	Username: "joko",
	Email:    "joko@gmail.com",
	Password: "123456",
	Address:  "Jakarta",
}

func TestUserRepsitoryTest_ValidateUsername(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success validate username", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)
		err := repo.ValidateUsernameAvailable(user.Username)

		assert.NoError(t, err)
	})

	t.Run("should error validate username", func(t *testing.T) {
		dbConn.Exec("INSERT INTO users (username, email, password, address) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Address)

		repo := repository.NewUserRepository(dbConn)
		err := repo.ValidateUsernameAvailable("joko")

		assert.Error(t, err)
	})
}

func TestUserRepsitoryTest_ValidateEmail(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success validate email", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)
		err := repo.ValidateEmailAvailable(user.Email)

		assert.NoError(t, err)
	})

	t.Run("should error validate email", func(t *testing.T) {
		dbConn.Exec("INSERT INTO users (username, email, password, address) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Password, user.Address)

		repo := repository.NewUserRepository(dbConn)
		err := repo.ValidateEmailAvailable(user.Email)

		assert.Error(t, err)
	})
}

func TestUserRepsitoryTest_Add(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success add user", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)
		err := repo.AddNewUser(user)

		assert.NoError(t, err)

		row := dbConn.QueryRow("SELECT id, username, email, password, address FROM users WHERE username = $1", user.Username)

		var result domain.User
		err = row.Scan(&result.Id, &result.Username, &result.Email, &result.Password, &result.Address)

		assert.NoError(t, err)

		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Password, result.Password)
		assert.Equal(t, user.Address, result.Address)
	})

	t.Run("should error add user with duplicate username", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		err := repo.AddNewUser(user)

		assert.Error(t, err)
	})
}

func TestUserRepsitoryTest_GetUserById(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should error get user by notfound id", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		_, err := repo.GetUserById(0)

		assert.Error(t, err)
	})

	t.Run("should success get user by id", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		repo.AddNewUser(user)

		userId := 0
		dbConn.QueryRow("SELECT id FROM users WHERE username = $1 LIMIT 1", user.Username).Scan(&userId)

		result, err := repo.GetUserById(userId)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Email, result.Email)
		assert.Empty(t, result.Password) // password should be empty
		assert.Equal(t, user.Address, result.Address)
	})
}

func TestUserRepsitoryTest_GetAllUsers(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success get all users length 1", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		repo.AddNewUser(user)

		result, err := repo.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("should success get all users length 2", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		user2 := domain.User{
			Username: "joni",
			Email:    "joni@gmail.com",
			Password: "123456",
			Address:  "Jakarta",
		}

		repo.AddNewUser(user2)

		result, err := repo.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, 2, len(result))
	})
}

func TestUserRepsitoryTest_UpdateUser(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success update user", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		repo.AddNewUser(user)

		dbConn.QueryRow("SELECT id FROM users WHERE username = $1 LIMIT 1", user.Username).Scan(&user.Id)

		user.Username = "joni"
		user.Email = "joni@gmail.com"

		err := repo.UpdateUserById(user)

		assert.NoError(t, err)
	})
}

func TestUserRepsitoryTest_DeleteUser(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should success delete user", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		repo.AddNewUser(user)

		dbConn.QueryRow("SELECT id FROM users WHERE username = $1 LIMIT 1", user.Username).Scan(&user.Id)

		err := repo.DeleteUserById(user.Id)
		assert.NoError(t, err)

		_, err = repo.GetUserById(user.Id)
		assert.Error(t, err)
	})
}

func TestUserRepsitoryTest_GetUserByUsername(t *testing.T) {
	defer truncateUsersTable()

	t.Run("should error get user by notfound username", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		_, err := repo.GetUserByUsername("joko")

		assert.Error(t, err)
	})

	t.Run("should success get user by username", func(t *testing.T) {
		repo := repository.NewUserRepository(dbConn)

		repo.AddNewUser(user)

		result, err := repo.GetUserByUsername(user.Username)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, result.Username)
		assert.Equal(t, user.Password, result.Password)
	})
}
