package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandlePinChange - не поддерживается для виртуальных карт
func (h *Handler) HandlePinChange(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
