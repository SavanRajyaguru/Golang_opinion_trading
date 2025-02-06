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
