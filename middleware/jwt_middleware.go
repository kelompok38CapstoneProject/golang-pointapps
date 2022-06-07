package middleware

import (
	"point/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int, nama string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["nama"] = nama
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT))
}
