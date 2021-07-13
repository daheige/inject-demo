package persistence

import (
	"database/sql"

	"github.com/daheige/inject-demo/internal/domain/repo"
)

// userRepository user repo
type userRepository struct {
	db *sql.DB
}

// NewUserRepository new user repo
func NewUserRepository(db *sql.DB) repo.UserRepository {
	return &userRepository{db: db}
}

// FindAll find user info
func (userRepo *userRepository) FindAll() ([]repo.User, error) {
	rows, err := userRepo.db.Query(`SELECT id, name, age FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]repo.User, 0, 10)
	for rows.Next() {
		user := repo.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
