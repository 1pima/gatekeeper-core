package transport

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
	processingv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/processing/v1"
)

func (c *GrpcClient) AdviceTransaction(
	ctx context.Context,
	reference string,
	extId string,
	trType iso8583.ProcessingCode,
	transactionAmount int64,
	transactionCurrency int32,
	settlementAmount int64,
	settlementCurrency int32,
	tokenId int64,
	mcc int64,
	bankData string,
	responseCode iso8583.ResponseCode,
) iso8583.ResponseCode {
	rq := &processingv1.AdviceTransactionRequest{
		Reference:         reference,
		ExtId:             extId,
		TrType:            string(trType),
		TransactionAmount: &processingv1.AmountData{Amount: transactionAmount, CurrencyCode: transactionCurrency},
		SettlementAmount:  &processingv1.AmountData{Amount: settlementAmount, CurrencyCode: settlementCurrency},
		TokenId:           tokenId,
		MccId:             mcc,
		BankData:          bankData,
		MpsCode:           string(responseCode),
	}

	rs := &processingv1.AdviceTransactionResponse{}
	rs, err := c.processingConn.AdviceTransaction(ctx, rq)
	if err != nil {
		logging.Log.Error("advice error: " + err.Error())
		return iso8583.SystemMalfunction
	}
	logging.Log.Info("advice success: " + rs.MpsCode)
	return iso8583.ResponseCode(rs.MpsCode)
}
