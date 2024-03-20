package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"try-standard-layout/internal/postgres"
	"try-standard-layout/internal/server/response"
)

type TestApi interface {
	TestGet(c echo.Context) error
	TestPost(c echo.Context) error
	TestPut(c echo.Context) error
}

type testApi struct {
	psg *postgres.Postgres
}

func (s *testApi) TestGet(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := s.psg.TestApi.List(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response.ToTestApis(res))
}

func (s *testApi) TestPost(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello Post")
}

func (s *testApi) TestPut(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello Put")
}
