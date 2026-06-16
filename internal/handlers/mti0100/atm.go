package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleATM - снятие через банкоматы, не поддерживает и ломает учет
func (h *Handler) HandleATM(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
