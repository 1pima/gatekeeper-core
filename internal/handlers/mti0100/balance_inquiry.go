package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleBalanceInquiry - запрос баланса (в рамках приложения не взаимодействуем с балансом)
func (h *Handler) HandleBalanceInquiry(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
