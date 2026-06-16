package mti0100

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
)

func (h *Handler) HandleServicePurchase(ctx context.Context, msg *iso8583.Message) *iso8583.Message {
	tokenId, err := h.client.GetTokenByHash(ctx, msg.CardHash())
	if err != nil {
		msg.ResponseCode = iso8583.SystemMalfunction
		return msg
	} else if tokenId == 0 {
		msg.ResponseCode = iso8583.NoCardRecord
		return msg
	}

	transactionId, err := h.client.Create(
		ctx,
		msg.RetrievalReferenceNumber,
		msg.STAN,
		msg.ProcessingCode, // payment GRPC enum
		msg.Amount,
		int32(msg.TransactionCurrencyCode),
		msg.SettlementAmount,
		int32(msg.SettlementCurrencyCode),
		tokenId,
		msg.MCC,
		msg.AdditionalData,
	)
	if err != nil {
		msg.ResponseCode = iso8583.SystemMalfunction
		return msg
	}

	msg.ResponseCode = h.client.Authorize(ctx, transactionId)
	return msg
}
