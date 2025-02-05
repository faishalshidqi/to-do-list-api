package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type refreshAuthnUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func (ru *refreshAuthnUsecase) GetUserByID(c context.Context, id primitive.ObjectID) (domains.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.userRepository.GetByID(ctx, id)
}

func (ru *refreshAuthnUsecase) CreateAccessToken(user domains.User, secret string, expiry int) (accessToken string, err error) {
	return tokenize.MakeJWT(user, secret, time.Duration(expiry)*time.Hour)
}

func (ru *refreshAuthnUsecase) CreateRefreshToken(user domains.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenize.MakeJWT(user, secret, time.Duration(expiry)*time.Hour)
}

func (ru *refreshAuthnUsecase) ExtractIDFromToken(token string, secret string) (primitive.ObjectID, error) {
	id, err := tokenize.ValidateJWT(token, secret)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}

func NewRefreshAuthnUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.RefreshAuthnUsecase {
	return &refreshAuthnUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
