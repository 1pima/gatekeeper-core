package router

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	"sync"
)

type Handler func(ctx context.Context, msg *iso8583.Message) *iso8583.Message

type Router struct {
	mu     sync.RWMutex
	routes map[iso8583.MTI]map[iso8583.ProcessingCode]Handler
}

func New() *Router {
	return &Router{routes: make(map[iso8583.MTI]map[iso8583.ProcessingCode]Handler)}
}

// Register биндит обработчик на конкретный MTI и Processing Code
func (r *Router) Register(mti iso8583.MTI, procCode iso8583.ProcessingCode, handler Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.routes[mti]; !exists {
		r.routes[mti] = make(map[iso8583.ProcessingCode]Handler)
	}
	r.routes[mti][procCode] = handler
}
