package repository

import (
	"fmt"
	"server/internal/common/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

const usersTable = "users"

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(user *entity.User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, firstname, lastname, username, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at, role", usersTable)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return r.db.QueryRow(query, id, user.Firstname, user.Lastname, user.Username, user.Email, user.Password).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Role)
}

func (r *UserPostgres) GetUserByUsernameAndPassword(username string, password string) (*entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username ILIKE $1 and password=$2 LIMIT 1", usersTable)
	err := r.db.Get(&user, query, username, password)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *UserPostgres) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return &user, err
}

func (r *UserPostgres) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username ILIKE $1", usersTable)
	err := r.db.Get(&user, query, username)

	return &user, err
}
