package repository

import (
	"server/internal/common/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByUsernameAndPassword(username string, password string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
}

type IpObject interface {
	CreateIpObject(ipObject *entity.IpObject) error
	GetIpObjectsByUserId(userId string) ([]*entity.IpObject, error)
	GetIpObjectsById(id string) (*entity.IpObject, error)
}

type Repository struct {
	UserRepository
	IpObject
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
		IpObject:       NewIpObjectPostgres(db),
	}
}
