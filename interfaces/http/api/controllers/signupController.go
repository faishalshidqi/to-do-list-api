package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"todo-list-api/commons/bootstrap"
	"todo-list-api/domains"
)

type SignupController struct {
	SignupUsecase domains.UserSignUpUsecase
	Env           *bootstrap.Env
}

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
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{
			Message: err.Error(),
		})
	}
	user := domains.User{
		ID:        primitive.NewObjectID(),
		Email:     request.Email,
		Password:  string(encryptedPassword),
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
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
