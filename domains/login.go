package domains

import (
	"context"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

/*
24

	7

-- x
28
14
-- +
168
*/
type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
