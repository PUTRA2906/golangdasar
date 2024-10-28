package repository

import (
	"database/sql"
	"myapp/entity"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.NAME, &user.AGE); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.NAME, &user.AGE)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *entity.User) error {
	result, err := r.db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.NAME, user.AGE)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	return nil
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
