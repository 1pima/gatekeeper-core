package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleMiniStatement - мини-выписка по карте, не поддерживается, баланс единый у корп карты для виртуальных карт
func (h *Handler) HandleMiniStatement(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
