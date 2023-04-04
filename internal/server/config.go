package server

import "net/http"

type Config struct {
	Addr string
	Port string
}

type Server struct {
	Params Config
	Router *http.ServeMux
}

func NewConfig() *Server {
	return &Server{
		Params: Config{
			Addr: "0.0.0.0", Port: "8080",
		},
		Router: http.NewServeMux(),
	}
}
