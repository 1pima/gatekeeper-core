package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleCardControl - управление картой (в рамках приложения не поддерживается, на стандии mvp отложено)
func (h *Handler) HandleCardControl(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
