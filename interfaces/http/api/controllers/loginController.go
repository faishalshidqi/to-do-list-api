package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type LoginController struct {
	LoginUsecase domains.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	loginRequest := domains.LoginRequest{}
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := lc.LoginUsecase.GetUserByEmail(c, loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{Message: err.Error()})
		return
	}
	err = tokenize.CheckPasswordHash(loginRequest.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{Message: "Invalid credentials"})
		return
	}
	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenKey, int(lc.Env.AccessTokenExpirationInHour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenKey, int(lc.Env.RefreshTokenExpirationInHour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domains.LoginResponse{
		Message:      "Successfully logged in",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
