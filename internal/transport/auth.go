package transport

import (
	"context"
	"gatekeeper-core/pkg/iso8583"
	processingv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/processing/v1"
)

func (c *GrpcClient) Authorize(ctx context.Context, transactionId string) iso8583.ResponseCode {
	rq := processingv1.AuthorizeTransactionRequest{TransactionId: transactionId}

	rs := &processingv1.AuthorizeTransactionResponse{}
	rs, err := c.processingConn.AuthorizeTransaction(ctx, &rq)
	if err != nil {
		return iso8583.SystemMalfunction
	}

	return iso8583.ResponseCode(rs.MpsCode)
}
