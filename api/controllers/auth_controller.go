package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	ctx := c.Request.Context()
	c.JSON(http.StatusOK, ctx)
}
