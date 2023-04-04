package server

import (
	"log"
	"net/http"
)

func (s *Server) Run() {
	mux := s.Router
	mux.HandleFunc("/", Index)

	bindAddr := s.Params.Addr + ":" + s.Params.Port
	srv := &http.Server{
		Addr:    bindAddr,
		Handler: middleware(mux.ServeHTTP),
	}
	log.Println("Server listen on", bindAddr)
	log.Fatal(srv.ListenAndServe())
}
