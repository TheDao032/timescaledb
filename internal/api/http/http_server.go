package http

import (
	"github.com/TheDao032/golang-architectures-demo/config"
	v1 "github.com/TheDao032/golang-architectures-demo/internal/api/http/v1"

	httpserver "github.com/TheDao032/go-backend-utils-architecture/http/server"
	"github.com/TheDao032/go-backend-utils-architecture/logger"
)

type Server struct {
	logger             logger.Logger
	cfg                *config.AppConfig
	healthcheckHandler *HealthcheckHandler
	acqHandler         *v1.ACQHandler
	navHandler         *v1.NAVHandler
	podHandler         *v1.PODHandler
	rawHandler         *v1.RAWHandler
	staHandler         *v1.STAHandler
}

func NewServer(
	logger logger.Logger,
	cfg *config.AppConfig,
	healthcheckHandler *HealthcheckHandler,
	acqHandler *v1.ACQHandler,
	navHandler *v1.NAVHandler,
	podHandler *v1.PODHandler,
	rawHandler *v1.RAWHandler,
	staHandler *v1.STAHandler,
) *Server {
  return &Server{logger, cfg, healthcheckHandler, acqHandler, navHandler, podHandler, rawHandler, staHandler}
}

func (s *Server) Run() {
	config := &httpserver.HttpServerConfig{
		Port:            s.cfg.Http.Port,
		Development:     s.cfg.Http.Development,
		ShutdownTimeout: s.cfg.Http.ShutdownTimeout,
		Resources:       s.cfg.Http.Resources,
		RateLimiting: &httpserver.RateLimitingConfig{
			RateFormat: s.cfg.Http.RateLimiting.RateFormat,
		},
	}
	httpServer, router := httpserver.NewServer(s.logger, *config)
	// In the future, if we have v2, v3..., we will add at here
	MapRoutes(router, s.healthcheckHandler, s.acqHandler, s.navHandler, s.podHandler, s.rawHandler, s.staHandler)
	httpServer.Run()
}
