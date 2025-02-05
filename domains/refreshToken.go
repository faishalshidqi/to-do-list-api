package domains

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RefreshTokenCollectionName = "refresh-token"
)

type RefreshAuthnRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
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
	GetUserByID(c context.Context, id primitive.ObjectID) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(token string, secret string) (primitive.ObjectID, error)
}

type RefreshAuthnRepository interface {
	Add(c context.Context, refreshRequest RefreshAuthnRequest) error
	Fetch(c context.Context, token string) (RefreshAuthnRequest, error)
	DeleteToken(c context.Context, token string) error
}
