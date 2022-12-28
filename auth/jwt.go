package auth

import (
	"myapp/customError"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

type CustomClaim struct {
	Username string `json:"username"`
}

func ValidateJWT(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessToken), nil
	})

	if token == nil {
		return "", customError.ErrInvalidParams(nil)
	}
	validKey := token.Claims
	claim := CustomClaim{}
	err = mapstructure.Decode(validKey, &claim)
	if err != nil {
		return "", err
	}
	return claim.Username, nil
}

func Generate_JWT_Token(username string) (string, error) {
	claims := jwt.MapClaims{"username": username}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
