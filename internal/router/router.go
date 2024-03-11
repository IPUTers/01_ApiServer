package router

import (
	"github.com/labstack/echo/v4"
	"try-standard-layout/internal/server"
)

type router struct {
	e   *echo.Echo
	rdb *server.Server
}

func New(
	e *echo.Echo,
	rdb *server.Server,
) {
	r := &router{e: e, rdb: rdb}
	api := e.Group("/api")
	r.group(api)
}
