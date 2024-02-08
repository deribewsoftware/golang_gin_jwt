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

	router.Run(":" + port)
	fmt.Println("server started at localhost:" + port)

}
