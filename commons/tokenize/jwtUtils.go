package tokenize

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
	"todo-list-api/domains"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetBearerToken(header http.Header) (string, error) {
	headers := header.Get("Authorization")
	if len(strings.Split(headers, " ")) < 2 {
		return "", errors.New("authorization header format must be Bearer {token}")
	}
	return strings.Split(headers, " ")[1], nil
}

func MakeJWT(user domains.User, tokenSecret string, expiresIn time.Duration) (string, error) {
	userID := user.ID
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ToDoList",
			Subject:   userID.Hex(),
		},
	)
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (primitive.ObjectID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return primitive.NilObjectID, err
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return primitive.NilObjectID, err
	}
	expirationTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return primitive.NilObjectID, err
	}
	issuedAt, err := token.Claims.GetIssuedAt()
	if err != nil {
		return primitive.NilObjectID, err
	}
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   subject,
		ExpiresAt: expirationTime,
		IssuedAt:  issuedAt,
	}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	if !parsedToken.Valid {
		return primitive.NilObjectID, errors.New("invalid token")
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return primitive.NilObjectID, errors.New("token expired")
	}
	parsedObjectID, err := primitive.ObjectIDFromHex(claims.Subject)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return parsedObjectID, nil
}
