package routes

import (
	"github.com/deribewsoftware/go_jwt/controllers"
	"github.com/gin-gonic/gin"
)

func adminRoutes(router *gin.Engine) {

	router.POST("/user/signup", controllers.Signup())
	router.POST("/user/login", controllers.Signin())

}
