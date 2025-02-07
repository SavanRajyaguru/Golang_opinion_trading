package mongoModels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a MongoDB user schema in Go
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // Auto-generated ObjectID
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Age       int                `bson:"age"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"update_at,omitempty"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"_id"` // Auto-generated ObjectID
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Age       int                `json:"age"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"update_at"`
}

func ConvertUserToResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
