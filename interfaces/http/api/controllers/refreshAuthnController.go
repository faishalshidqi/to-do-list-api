package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/domains"
)

type RefreshAuthnController struct {
	RefreshAuthnUsecase domains.RefreshAuthnUsecase
	RefreshAuthnRepo    domains.RefreshAuthnRepository
	Env                 *bootstrap.Env
}

// RefreshToken Refresh Authentication, generating new access and refresh token godoc
//
//	@Summary		Refresh Authentication
//	@Description	Generating new access token using a refresh token. Only valid refresh token will generate new
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	body		string	true	"refresh token possessed by the user"
//	@Success		200				{object}	domains.RefreshAuthnResponse
//	@Failure		400				{object}	domains.ErrorResponse
//	@Failure		401				{object}	domains.ErrorResponse
//	@Failure		500				{object}	domains.ErrorResponse
//	@Router			/api/auth [put]
func (rc *RefreshAuthnController) RefreshToken(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
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
	err = rc.RefreshAuthnRepo.Add(c, refreshAuthnRequest)
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
