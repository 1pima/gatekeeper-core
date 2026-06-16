package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

// HandleCashin - внесение наличных через ATM (не поддерживается, ломает бизнес логику нарушая правила учета расходов)
func (h *Handler) HandleCashin(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	msg.ResponseCode = iso8583.NoActionTaken
	return msg
}
