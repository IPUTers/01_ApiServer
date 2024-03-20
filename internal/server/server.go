package server

import (
	"try-standard-layout/internal/postgres"
)

type Server struct {
	TestApi *testApi
}

func New(
	psg *postgres.Postgres,
) *Server {
	return &Server{
		TestApi: &testApi{psg},
	}
}
