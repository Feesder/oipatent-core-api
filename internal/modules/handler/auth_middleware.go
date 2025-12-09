package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	op := "handler.auth_middleware.userIdentity"

	log := logrus.WithField("op", op)

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Error("empty auth header")
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		log.Error("invalid auth header")
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	claims, err := h.Services.ParseAccessToken(headerParts[1])
	if err != nil {
		log.Error(err.Error())
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, claims.UserId)
}
