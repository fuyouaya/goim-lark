package server

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/router"
	"lark/pkg/common/xgin"
)

type server struct {
	cfg       *config.Config
	ginServer *xgin.GinServer
}

func NewServer() (s *server) {
	return &server{
		cfg:       config.GetConfig(),
		ginServer: xgin.NewGinServer(),
	}
}

func (s *server) Run() {
	router.Register(s.ginServer.Engine)
	s.ginServer.Run(s.cfg.Port)
}
