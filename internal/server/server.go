package server

type Server struct {
	TestApi *testApi
}

func New() *Server {
	return &Server{
		TestApi: &testApi{},
	}
}
