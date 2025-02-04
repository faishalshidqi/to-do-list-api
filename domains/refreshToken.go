package domains

import "context"

type RefreshAuthnRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshAuthnData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshAuthnResponse struct {
	Message string           `json:"message"`
	Data    RefreshAuthnData `json:"data"`
}

type RefreshAuthnUsecase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(token string, secret string) (string, error)
}
