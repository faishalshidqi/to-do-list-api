package domains

import "context"

const (
	RefreshTokenCollectionName = "refresh-token"
)

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

type RefreshAuthenticationRepository interface {
	Add(c context.Context, token string) error
	Fetch(c context.Context, token string) (string, error)
	DeleteToken(c context.Context, token string) error
}
