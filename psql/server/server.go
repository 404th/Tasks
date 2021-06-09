package main

import "net/http"

type Server struct {
	httpServer http.Server
}

func (s *Server) Run(port string) {

}
