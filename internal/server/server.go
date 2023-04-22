package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"mal-rest-api/internal/config"
	"mal-rest-api/pkg/logger"
)

type Server struct {
	router *httprouter.Router
	logger *logger.Logger
}

func NewServer(router *httprouter.Router, logger *logger.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

func (s *Server) Run() {
	logger := logger.GetLogger()
	cfg := config.GetConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      s.router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	logger.Fatal().Err(server.Serve(listener))
}
