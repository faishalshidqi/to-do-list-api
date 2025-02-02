package domains

import "context"

type UserSignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserSignUpUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *User, secret string, expiry int) (string, error)
}
