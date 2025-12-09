package handler

import (
	"server/internal/modules/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Services  *service.Service
	Validator *validator.Validate
}

type Deps struct {
	Services  *service.Service
	Validator *validator.Validate
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		Services:  deps.Services,
		Validator: deps.Validator,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin", "Accept"},
		ExposeHeaders:    []string{"Set-Cookie", "Content-Length"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/refresh", h.Refresh)
	}

	profile := router.Group("/")
	profile.Use(h.UserIdentity)
	profile.GET("/me", h.me)

	ipObject := router.Group("/ip-objects")
	{
		ipObject.Use(h.UserIdentity)
		ipObject.POST("", h.createIpObject)
		ipObject.GET("/user/:id", h.getIpObjectByUserId)
	}

	return router
}
