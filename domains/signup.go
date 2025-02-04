package domains

import "context"

type UserSignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserSignUpUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
}
