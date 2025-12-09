package service

import (
	"errors"
	"net/http"
	"server/internal/common/dto"
	"server/internal/config"
	"server/internal/modules/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationService struct {
	repos *repository.Repository
	cfg   *config.Config
	v     *validator.Validate
}

func NewAuthorizationService(repository *repository.Repository, cfg *config.Config, v *validator.Validate) *AuthorizationService {
	return &AuthorizationService{
		repos: repository,
		cfg:   cfg,
		v:     v,
	}
}

type AccessClaims struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func (s *AuthorizationService) GenerateToken(username string, password string) (*dto.TokensDto, error) {
	generatePassword(&password, s.cfg.Security.Salt)
	user, err := s.repos.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	accessToken, aexp, err := s.SignAccess(user.Id, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, rexp, err := s.SignRefresh(user.Id)
	if err != nil {
		return nil, err
	}

	return &dto.TokensDto{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  aexp,
		RefreshExpires: rexp,
	}, nil
}

func (s *AuthorizationService) GenerateTokenByUser(user *dto.UserDto) (*dto.TokensDto, error) {
	accessToken, aexp, err := s.SignAccess(user.Id, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, rexp, err := s.SignRefresh(user.Id)
	if err != nil {
		return nil, err
	}

	return &dto.TokensDto{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  aexp,
		RefreshExpires: rexp,
	}, nil
}

func (s *AuthorizationService) SignAccess(userId string, role string) (string, time.Time, error) {
	now := time.Now()
	exp := time.Now().Add(s.cfg.Token.AccessTTL)
	claims := AccessClaims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    s.cfg.Token.Issuer,
			Subject:   "access",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := t.SignedString([]byte(s.cfg.Token.AccessSecret))
	return str, exp, err
}

func (s *AuthorizationService) SignRefresh(userId string) (string, time.Time, error) {
	now := time.Now()
	exp := time.Now().Add(s.cfg.Token.RefreshTTL)
	claims := RefreshClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    s.cfg.Token.Issuer,
			Subject:   "refresh",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := t.SignedString([]byte(s.cfg.Token.RefreshSecret))
	return str, exp, err
}

func (s *AuthorizationService) ParseAccessToken(accessToken string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok {
		return nil, errors.New("token claims are not type AccessClaims")
	}

	return claims, nil
}

func (s *AuthorizationService) ParseRefreshToken(refreshToken string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok {
		return nil, errors.New("token claims are not type AccessClaims")
	}

	return claims, nil
}

func (s *AuthorizationService) SetCookie(c *gin.Context, name, value string, maxAge int) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		name,
		value,
		maxAge,
		s.cfg.Cookie.Path,
		s.cfg.Cookie.Domain,
		s.cfg.Cookie.Secure,
		s.cfg.Cookie.HttpOnly,
	)
}

func (s *AuthorizationService) GetCookie(c *gin.Context, name string) (string, error) {
	token, err := c.Cookie(name)
	if err != nil {
		return "", err
	}

	return token, nil
}
