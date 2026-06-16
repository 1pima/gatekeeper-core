package main

import (
	"context"
	"gatekeeper-core/internal/handlers/mti0100"
	"gatekeeper-core/internal/handlers/mti0120"
	"gatekeeper-core/internal/handlers/mti0800"
	"gatekeeper-core/internal/router"
	"gatekeeper-core/internal/transport"
	"gatekeeper-core/pkg/iso8583"
)

func isoRouter() *router.Router {
	r := router.New()

	h0800 := mti0800.NewHandler()
	r.Register(iso8583.MTINetworkManagementRequest, iso8583.ServicePurchase, h0800.HandleEcho)
	r.Register(iso8583.MTINetworkAdviceRequest, iso8583.ServicePurchase, h0800.HandleEcho)

	h0100 := mti0100.New(transport.GatekeeperEngine)
	authorizationHandlers := map[iso8583.ProcessingCode]func(context.Context, *iso8583.Message) *iso8583.Message{
		iso8583.ServicePurchase:     h0100.HandleServicePurchase,
		iso8583.ATMWithdrawal:       h0100.HandleATM,
		iso8583.AccountFunding:      h0100.HandleAccountFunding,
		iso8583.Unique:              h0100.HandleUnique,
		iso8583.CashAdvance:         h0100.HandleCashAdvance,
		iso8583.Payment:             h0100.HandlePayment,
		iso8583.CashIn:              h0100.HandleCashin,
		iso8583.BalanceInquiry:      h0100.HandleBalanceInquiry,
		iso8583.MiniStatement:       h0100.HandleMiniStatement,
		iso8583.AccountVerification: h0100.HandleAccountVerification,
		iso8583.CardControl:         h0100.HandleCardControl,
		iso8583.PinChange:           h0100.HandlePinChange,
	}
	for opCode, handler := range authorizationHandlers {
		r.Register(iso8583.MTIAuthorizationRequest, opCode, handler)
	}

	h0120 := mti0120.New(transport.GatekeeperEngine)
	authorizationAdviceHandlers := map[iso8583.ProcessingCode]func(context.Context, *iso8583.Message) *iso8583.Message{
		iso8583.ServicePurchase:     h0120.HandleServicePurchase,
		iso8583.ATMWithdrawal:       h0120.HandleATM,
		iso8583.AccountFunding:      h0120.HandleAccountFunding,
		iso8583.Unique:              h0120.HandleUnique,
		iso8583.CashAdvance:         h0120.HandleCashAdvance,
		iso8583.Payment:             h0120.HandlePayment,
		iso8583.CashIn:              h0120.HandleCashin,
		iso8583.BalanceInquiry:      h0120.HandleBalanceInquiry,
		iso8583.MiniStatement:       h0120.HandleMiniStatement,
		iso8583.AccountVerification: h0120.HandleAccountVerification,
		iso8583.CardControl:         h0120.HandleCardControl,
		iso8583.PinChange:           h0120.HandlePinChange,
	}
	for opCode, handler := range authorizationAdviceHandlers {
		r.Register(iso8583.MTIAuthorizationAdviceRequest, opCode, handler)
	}

	return r
}
