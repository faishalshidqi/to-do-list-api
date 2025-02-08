package domains

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	UserCollectionName = "users"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"createdAt" time_format:"01-01-2001 15:04:05"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updatedAt" time_format:"01-01-2001 15:04:05"`
}

type UserRepository interface {
	Add(ctx context.Context, user *User) error
	Fetch(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (User, error)
}
