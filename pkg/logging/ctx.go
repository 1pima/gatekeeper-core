package logging

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	"log/slog"
)

type ctxKey string

// Ключи контекста для логирования
const (
	MTI                     ctxKey = "mti"
	ProcessingCode          ctxKey = "processing_code"
	STAN                    ctxKey = "stan"
	Amount                  ctxKey = "amount"
	SettlementAmount        ctxKey = "settlement_amount"
	TransactionCurrencyCode ctxKey = "currency"
	SettlementCurrencyCode  ctxKey = "settlement_currency"
	MCC                     ctxKey = "mcc"
)

// ContextHandler оборачивает стандартный slog.Handler
type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if mti, ok := ctx.Value(MTI).(string); ok {
		r.AddAttrs(slog.String("mti", mti))
	}
	if processingCode, ok := ctx.Value(ProcessingCode).(string); ok {
		r.AddAttrs(slog.String("processing_code", processingCode))
	}
	if stan, ok := ctx.Value(STAN).(string); ok {
		r.AddAttrs(slog.String("stan", stan))
	}
	if amount, ok := ctx.Value(Amount).(string); ok {
		r.AddAttrs(slog.String("amount", amount))
	}
	if settlementAmount, ok := ctx.Value(SettlementAmount).(string); ok {
		r.AddAttrs(slog.String("settlement_amount", settlementAmount))
	}
	if currency, ok := ctx.Value(TransactionCurrencyCode).(string); ok {
		r.AddAttrs(slog.String("currency", currency))
	}
	if settlementCurrency, ok := ctx.Value(SettlementCurrencyCode).(string); ok {
		r.AddAttrs(slog.String("settlement_currency", settlementCurrency))
	}
	if mcc, ok := ctx.Value(MCC).(string); ok {
		r.AddAttrs(slog.String("mcc", mcc))
	}

	return h.Handler.Handle(ctx, r)
}

func WithMessage(ctx context.Context, msg iso8583.Message) context.Context {
	ctx = context.WithValue(ctx, MTI, msg.MTI)
	ctx = context.WithValue(ctx, ProcessingCode, msg.ProcessingCode)
	ctx = context.WithValue(ctx, STAN, msg.STAN)
	ctx = context.WithValue(ctx, Amount, msg.Amount)
	ctx = context.WithValue(ctx, SettlementAmount, msg.SettlementAmount)
	ctx = context.WithValue(ctx, TransactionCurrencyCode, msg.TransactionCurrencyCode)
	ctx = context.WithValue(ctx, SettlementCurrencyCode, msg.SettlementCurrencyCode)
	ctx = context.WithValue(ctx, MCC, msg.MCC)
	return ctx
}
