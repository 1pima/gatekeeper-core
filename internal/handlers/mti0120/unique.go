package mti0120

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

func (h *Handler) HandleUnique(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	return h.baseAdviceHandler(ctx, msg)
}
