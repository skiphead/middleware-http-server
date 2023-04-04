package server

import (
	"log"
	"net/http"
)

func (s *Server) Run() {

	mux := s.Router
	mux.HandleFunc("/", Index)

	srv := &http.Server{
		Addr:    s.Params.BindAddr,
		Handler: middleware(mux.ServeHTTP),
	}
	log.Println("Server listen on", s.Params.BindAddr)
	log.Fatal(srv.ListenAndServe())

}
