package config

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTConfig struct {
	// JWT signature secret
	JWTSecret string
}

func NewJWTConfig(secret string) *JWTConfig {
	return &JWTConfig{
		JWTSecret: secret,
	}
}

type JWTCustomClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (config *JWTConfig) GenerateToken(userId, role string, exp time.Time) (string, error) {
	claims := JWTCustomClaims{
		UserId: userId,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := config.JWTSecret

	return token.SignedString([]byte(jwtSecret))
}

func (config *JWTConfig) ExtractToken(c echo.Context) (*JWTCustomClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")

	sanitizedTokenBearer := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(sanitizedTokenBearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constant.ErrInvalidTokenHeader
		}

		jwtSecret := config.JWTSecret

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		customClaim := JWTCustomClaims{}

		claims := token.Claims.(jwt.MapClaims)
		customClaim.UserId = claims["user_id"].(string)
		customClaim.Role = claims["role"].(string)

		return &customClaim, nil
	}

	return nil, err
}
