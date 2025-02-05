package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type LogoutController struct {
	RefreshAuthnRepo domains.RefreshAuthnRepository
	Env              *bootstrap.Env
}

func (lc *LogoutController) Logout(c *gin.Context) {
	refreshAuthnRequest := domains.RefreshAuthnRequest{}
	err := c.ShouldBind(&refreshAuthnRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	_, err = tokenize.ValidateJWT(refreshAuthnRequest.RefreshToken, lc.Env.RefreshTokenKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	fetchResult, err := lc.RefreshAuthnRepo.Fetch(c, refreshAuthnRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{
			Message: "Invalid refresh token",
		})
		return
	}
	err = lc.RefreshAuthnRepo.DeleteToken(c, fetchResult.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domains.SuccessResponse{
		Message: "Logged out successfully",
	})
}
