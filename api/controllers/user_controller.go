package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	c.JSON(http.StatusOK, ctx)
}
