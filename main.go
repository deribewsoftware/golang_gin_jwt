package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.userRoutes(router)
	routes.adminRoutes(router)

	//test api
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "your api is successfully",
		})
	})

	router.Run(":" + port)
	fmt.Println("server started at localhost:" + port)

}
