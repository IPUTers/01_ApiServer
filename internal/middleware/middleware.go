package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Set(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Logger())
	e.Use(SetDB(db))
}

func SetDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "db", db)))
			return next(c)
		}
	}
}
