package mti0800

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// HandleEcho отвечает на пинг от банка (0800 -> 0810)
func (h *Handler) HandleEcho(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	logging.Log.Debug("Received 0800 Echo Request from Bank")

	msg.MTI = msg.MTI.ResponseMTI()
	msg.ResponseCode = iso8583.Approved
	return msg
}
