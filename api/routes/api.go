package routes

import (
	"github.com/gin-gonic/gin"
	"pxj/courseSystem/api/controllers"
)

var (
	Routes *gin.Engine
)

func InitRoutes() {
	Routes := gin.Default()
	Routes.POST("/login", controllers.Login)

	api := Routes.Group("/api")
	api.POST("/users/getUserById", controllers.GetUserById)
}
