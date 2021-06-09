package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handlers,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 22,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
