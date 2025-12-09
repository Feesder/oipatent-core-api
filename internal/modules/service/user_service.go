package service

import (
	"crypto/sha1"
	"fmt"
	"server/internal/common/dto"
	"server/internal/common/entity"
	"server/internal/common/mapper"
	"server/internal/config"
	"server/internal/modules/repository"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	repos *repository.Repository
	cfg   *config.Config
	v     *validator.Validate
}

func NewUserService(repository *repository.Repository, cfg *config.Config, v *validator.Validate) *UserService {
	return &UserService{
		repos: repository,
		cfg:   cfg,
		v:     v,
	}
}

func (s *UserService) CreateUser(user dto.SignUpDto) (*dto.UserDto, error) {
	generatePassword(&user.Password, s.cfg.Security.Salt)

	userEntity := &entity.User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
	}

	if err := s.repos.CreateUser(userEntity); err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDto(userEntity), nil
}

func (s *UserService) GetUserById(id string) (*dto.UserDto, error) {
	userEntity, err := s.repos.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDto(userEntity), nil
}

func (s *UserService) GetUserByUsername(username string) (*dto.UserDto, error) {
	userEntity, err := s.repos.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDto(userEntity), nil
}

func (s *UserService) GetUserByUsernameAndPassword(username string, password string) (*dto.UserDto, error) {
	generatePassword(&password, s.cfg.Security.Salt)
	userEntity, err := s.repos.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserEntityToUserDto(userEntity), nil
}

func generatePassword(password *string, salt string) {
	hash := sha1.New()
	hash.Write([]byte(*password))
	*password = fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
