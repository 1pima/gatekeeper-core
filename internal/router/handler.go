package router

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// Handle принимает сообщение и отправляет в нужный контроллер
func (r *Router) Handle(ctx context.Context, msg *iso8583.Message) *iso8583.Message { // todo там где nil - отдавать отмену с необходимым кодом
	r.mu.RLock()
	mtiRoutes, mtiExists := r.routes[msg.MTI]
	var handler Handler
	var handlerExists bool

	procCode := iso8583.ProcessingCode(msg.ProcessingCode)
	if mtiExists {
		handler, handlerExists = mtiRoutes[procCode]
	}

	r.mu.RUnlock()

	if !mtiExists || !handlerExists {
		return r.CreateRejectMessage(msg, msg.MTI, iso8583.FormatError)
	}

	return handler(ctx, msg)
}
