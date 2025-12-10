package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
