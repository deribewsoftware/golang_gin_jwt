package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required , min=2 , max=100"`
	LastName     *string            `json:"last_name" validate:"required , min=2 , max=100"`
	Phone        *string            `json:"phone"     validate:"required"`
	Email        *string            `json:"email"     validate:"required"`
	Password     *string            `json:"password"    validate:"required ,min=6"`
	UserType     *string            `json:"user_type" validate:"required ,eq=ADMIN|eq=USER"`
	Token        *string            `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
