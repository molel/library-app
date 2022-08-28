package server

import (
	"context"
	"library-app/entities"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(configs *entities.HTTPServerConfigs) error {
	s.httpServer = &http.Server{
		Addr:           configs.Addr,
		Handler:        configs.Handler,
		ReadTimeout:    configs.ReadTimeout,
		WriteTimeout:   configs.WriteTimeout,
		MaxHeaderBytes: configs.MaxHeaderBytes,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
