package usecase

import (
	"context"
	"time"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type loginUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func (lu loginUsecase) GetUserByEmail(c context.Context, email string) (domains.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu loginUsecase) CreateAccessToken(user *domains.User, secret string, expiry int) (accessToken string, err error) {
	return tokenize.MakeJWT(*user, secret, time.Duration(expiry)*time.Hour)
}

func (lu loginUsecase) CreateRefreshToken(user *domains.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenize.MakeJWT(*user, secret, time.Duration(expiry)*time.Hour)
}

func NewLoginUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
