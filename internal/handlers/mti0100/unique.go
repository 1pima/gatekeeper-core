package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleUnique - ломает и усложняет учет корп расходов, такие транзакции не обрабатываем
func (h *Handler) HandleUnique(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
