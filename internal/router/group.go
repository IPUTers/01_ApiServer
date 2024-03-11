package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *router) group(g *echo.Group) {
	g.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello echo")
	})

	test := g.Group("/test")
	{
		test.GET("", r.rdb.TestApi.TestGet)
		test.POST("", r.rdb.TestApi.TestPost)
		test.PUT("", r.rdb.TestApi.TestPut)
	}
}
