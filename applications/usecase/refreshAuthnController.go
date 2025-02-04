package usecase

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/domains"
)

type refreshAuthnController struct {
	RefreshAuthnUsecase domains.RefreshAuthnUsecase
	Env                 *bootstrap.Env
}

func (rc *refreshAuthnController) RefreshToken(c *gin.Context) {
	refreshAuthnRequest := domains.RefreshAuthnRequest{}
	err := c.ShouldBind(&refreshAuthnRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	id, err := rc.RefreshAuthnUsecase.ExtractIDFromToken(refreshAuthnRequest.RefreshToken, rc.Env.RefreshTokenKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	user, err := rc.RefreshAuthnUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{
			Message: "user not found",
		})
		return
	}
	accessToken, err := rc.RefreshAuthnUsecase.CreateAccessToken(user, rc.Env.AccessTokenKey, rc.Env.AccessTokenExpirationInHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	refreshToken, err := rc.RefreshAuthnUsecase.CreateRefreshToken(user, rc.Env.RefreshTokenKey, rc.Env.RefreshTokenExpirationInHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domains.RefreshAuthnResponse{
		Message: "Successfully refreshed access token",
		Data: domains.RefreshAuthnData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
