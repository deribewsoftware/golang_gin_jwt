package routes

import (
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine) {
	router.Get("/users", controllers.getUsers())
	router.Get("/user/:id", controllers.getUser())

}
