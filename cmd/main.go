package main

import (
	"try-standard-layout/internal/middleware"
	"try-standard-layout/internal/postgres"
	"try-standard-layout/internal/router"
	"try-standard-layout/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	psg := postgres.New()

	db, err := postgres.Connect()
	if err != nil {
		return
	}
	defer postgres.Close(db)

	middleware.Set(e, db)

	srv := server.New(psg)

	router.New(e, srv)

	e.Logger.Fatal(e.Start(":9090"))

}
