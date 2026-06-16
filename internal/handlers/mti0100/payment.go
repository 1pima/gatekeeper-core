package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandlePayment - пополнение карты при p2p, невозможно, пока не ясно как в таком случае будет садится баланс
func (h *Handler) HandlePayment(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
