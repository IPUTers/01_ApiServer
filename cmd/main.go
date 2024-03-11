package main

import (
	"try-standard-layout/internal/router"
	"try-standard-layout/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	srv := server.New()

	router.New(e, srv)

	e.Logger.Fatal(e.Start(":9090"))

}
