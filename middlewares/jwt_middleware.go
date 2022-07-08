package middlewares

import (
	"net/http"
	constant "point/constants"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int, nama, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["nama"] = nama
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT))
}
func GetToken(c echo.Context) string {
	header := c.Request().Header.Get("Authorization")
	token := strings.Split(header, " ")[1]
	return token
}

type Claims struct {
	jwt.StandardClaims
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
}

func ParseJWT(tokenstr string) (*Claims, error) {
	hmacSecret := constant.SECRET_JWT
	var claims *Claims

	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return claims, err
}

func AdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := GetToken(c)
		claims, _ := ParseJWT(token)
		if claims.Role != "admin" {
			return c.JSON(http.StatusForbidden, "access denied, not admin")
		}

		return next(c)
	}
}
