package handler

import (
	"fmt"
	"net/http"
	"server/internal/common/dto"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createIpObject(c *gin.Context) {
	op := "handler.ip_object.handler.createIpObject"

	log := logrus.WithField("op", op)

	var input dto.CreateIpObjectDto
	if err := c.BindJSON(&input); err != nil {
		log.Error(fmt.Sprintf("request body error: %s", err.Error()))
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Validator.Struct(input); err != nil {
		log.Error(fmt.Sprintf("validate error: %s", err.Error()))
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ipObject, err := h.Services.CreateIpObject(input)
	if err != nil {
		log.Error(fmt.Sprintf("create ip object error: %s", err.Error()))
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Info(fmt.Sprintf("ip object created by id: %s", ipObject.Id))

	c.JSON(http.StatusOK, gin.H{
		"data": ipObject,
	})
}

func (h *Handler) getIpObjectByUserId(c *gin.Context) {
	op := "handler.ip_object.handler.getIpObjectByUserId"

	log := logrus.WithField("op", op)

	userId := c.Param("id")

	ipObject, err := h.Services.GetIpObjectsByUserId(userId)
	if err != nil {
		log.Error(fmt.Sprintf("get ip object error: %s", err.Error()))
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ipObject,
	})
}
