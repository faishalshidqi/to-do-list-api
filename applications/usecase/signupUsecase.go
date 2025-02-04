package usecase

import (
	"context"
	"time"
	"todo-list-api/domains"
)

type signupUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func (su *signupUsecase) Create(c context.Context, user *domains.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Add(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domains.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func NewSignupUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.UserSignUpUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
