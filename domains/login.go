package domains

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(c context.Context, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(c context.Context, secret string, expiry int) (refreshToken string, err error)
}
