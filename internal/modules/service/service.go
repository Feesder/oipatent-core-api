package service

import (
	"server/internal/common/dto"
	"server/internal/config"
	"server/internal/modules/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Authorization interface {
	GenerateTokenByUser(user *dto.UserDto) (*dto.TokensDto, error)
	SetCookie(c *gin.Context, name, value string, maxAge int)
	GetCookie(c *gin.Context, name string) (string, error)
	ParseAccessToken(token string) (*AccessClaims, error)
	ParseRefreshToken(token string) (*RefreshClaims, error)
}

type User interface {
	CreateUser(user dto.SignUpDto) (*dto.UserDto, error)
	GetUserById(id string) (*dto.UserDto, error)
	GetUserByUsername(username string) (*dto.UserDto, error)
	GetUserByUsernameAndPassword(username string, password string) (*dto.UserDto, error)
}

type IpObject interface {
	CreateIpObject(ipObject dto.CreateIpObjectDto) (*dto.IpObjectDto, error)
	GetIpObjectsByUserId(userId string) ([]*dto.IpObjectDto, error)
	GetIpObjectsById(id string) (*dto.IpObjectDto, error)
}

type Service struct {
	Authorization
	User
	IpObject
}

type Deps struct {
	Repos     *repository.Repository
	Cfg       *config.Config
	Validator *validator.Validate
}

func NewService(deps *Deps) *Service {
	return &Service{
		Authorization: NewAuthorizationService(deps.Repos, deps.Cfg, deps.Validator),
		User:          NewUserService(deps.Repos, deps.Cfg, deps.Validator),
		IpObject:      NewIpOjbectService(deps.Repos, deps.Cfg),
	}
}
