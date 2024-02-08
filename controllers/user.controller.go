package controllers

import (
	"github/gin-gonic/gin"

	"github.com/deribewsoftware/go_jwt/database"
	"github.com/deribewsoftware/go_jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()
func VerifyPassword(password string)

func getUsers() *gin.HandlerFunc {
	var user *models.User

	return func(c *gin.Context) {
		c.Bind

	}

}

func getUserById() {}

func Signup() {}
func Signin() {}
