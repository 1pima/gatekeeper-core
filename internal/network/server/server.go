package server

import (
	"gatekeeper-core/internal/config"
	"gatekeeper-core/internal/router"
	"gatekeeper-core/pkg/iso8583"
)

type Server struct {
	address      string
	router       *router.Router
	outboundChan chan<- *iso8583.Message
}

func New(r *router.Router, outChan chan<- *iso8583.Message) *Server {
	return &Server{
		address:      config.Settings.ServerAddr,
		router:       r,
		outboundChan: outChan,
	}
}
