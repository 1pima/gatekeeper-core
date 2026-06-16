package transport

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	"gatekeeper-core/pkg/logging"
	processingv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/processing/v1"
)

func (c *GrpcClient) Create(
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
) (string, error) {
	rq := &processingv1.CreateTransactionRequest{
		Reference:         reference,
		ExtId:             extId,
		TrType:            string(trType),
		TransactionAmount: &processingv1.AmountData{Amount: transactionAmount, CurrencyCode: transactionCurrency},
		SettlementAmount:  &processingv1.AmountData{Amount: settlementAmount, CurrencyCode: settlementCurrency},
		TokenId:           tokenId,
		MccId:             mcc,
		BankData:          bankData,
	}

	rs := &processingv1.CreateTransactionResponse{}
	rs, err := c.processingConn.CreateTransaction(ctx, rq)
	if err != nil {
		return "", err
	}
	logging.Log.Info("successfully created transaction: " + rs.TransactionId)
	return rs.TransactionId, nil
}
