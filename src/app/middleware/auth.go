package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"net/http"
)

type AuthMiddleware struct {
	jwtConfig *config.JWTConfig
}

func NewAuthMiddleware(jwtConfig *config.JWTConfig) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: jwtConfig,
	}
}

// IS USER
func (m *AuthMiddleware) IsUserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		payload, err := m.jwtConfig.ExtractToken(ctx)

		if err != nil {
			return ctx.JSON(http.StatusUnauthorized,
				helper.MessageErrorResponse(err.Error()))
		}

		if payload.Role != "USER" {
			return ctx.JSON(http.StatusForbidden,
				helper.MessageErrorResponse(constant.ErrAccessForbidden.Error()))
		}

		return next(ctx)
	}
}

// IS ADMIN
func (m *AuthMiddleware) IsAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		payload, err := m.jwtConfig.ExtractToken(ctx)

		if err != nil {
			return ctx.JSON(http.StatusUnauthorized,
				helper.MessageErrorResponse(err.Error()))
		}

		if payload.Role != "ADMIN" {
			return ctx.JSON(http.StatusForbidden,
				helper.MessageErrorResponse(constant.ErrAccessForbidden.Error()))
		}

		return next(ctx)
	}
}

func (m *AuthMiddleware) IsAuthenticated() echo.MiddlewareFunc {
	return middleware.JWT([]byte(m.jwtConfig.JWTSecret))
}
