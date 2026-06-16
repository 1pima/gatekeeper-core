package mti0120

import (
	"context"
	"gatekeeper-core/internal/transport"
	"gatekeeper-core/pkg/iso8583"
)

type Handler struct {
	client transport.Client
}

func New(client transport.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) baseAdviceHandler(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	tokenId, err := h.client.GetTokenByHash(ctx, msg.CardHash())
	if err != nil {
		msg.ResponseCode = iso8583.SystemMalfunction
		return msg
	} else if tokenId == 0 {
		msg.ResponseCode = iso8583.NoCardRecord
		return msg
	}

	msg.ResponseCode = h.client.AdviceTransaction(
		ctx,
		msg.RetrievalReferenceNumber,
		msg.STAN,
		msg.ProcessingCode,
		msg.Amount,
		int32(msg.TransactionCurrencyCode),
		msg.SettlementAmount,
		int32(msg.SettlementCurrencyCode),
		tokenId,
		msg.MCC,
		msg.AdditionalData,
		msg.ResponseCode,
	)
	return msg
}
