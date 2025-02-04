package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/commons/tokenize"
	"todo-list-api/domains"
)

type SignupController struct {
	SignupUsecase domains.UserSignUpUsecase
	Env           *bootstrap.Env
}

// Signup AddUser godoc
//
//	@Summary		Register A User
//	@Description	New user must have a unique email address
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			email		body		string	true	"email address of the new user, must be unique"	Format(email)
//	@Param			password	body		string	true	"password of the new user"
//	@Param			name		body		string	true	"name of the new user"
//	@Success		201			{object}	domains.SuccessResponse
//	@Failure		409			{object}	domains.ErrorResponse
//	@Failure		500			{object}	domains.ErrorResponse
//	@Router			/api/auth/register [post]
func (sc *SignupController) Signup(c *gin.Context) {
	request := domains.UserSignUpRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	_, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domains.ErrorResponse{
			Message: "User with this email already exists",
		})
		return
	}
	encryptedPassword, err := tokenize.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	user := domains.User{
		ID:        primitive.NewObjectID(),
		Email:     request.Email,
		Password:  encryptedPassword,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, domains.SuccessResponse{
		Message: "User created successfully",
	})
}
