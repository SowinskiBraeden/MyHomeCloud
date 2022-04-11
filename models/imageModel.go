package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID          primitive.ObjectID `bson:"_id"`
	FSName      string             `json:"fsname" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Album       string             `json:"album" validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
