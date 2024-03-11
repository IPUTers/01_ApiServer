package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type TestApi interface {
	TestGet(c echo.Context) error
	TestPost(c echo.Context) error
	TestPut(c echo.Context) error
}

type testApi struct {
}

func (s *testApi) TestGet(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello Get")
}

func (s *testApi) TestPost(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello Post")
}

func (s *testApi) TestPut(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello Put")
}
