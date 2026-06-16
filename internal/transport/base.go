package transport

import (
	"context"
	"gatekeeper-core/internal/config"
	"gatekeeper-core/pkg/iso8583"

	// $ make generate-contracts
	processingv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/processing/v1"
	vaultv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/vault/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GatekeeperEngine Client

type Client interface {
	Create(ctx context.Context, reference string, extId string, trType iso8583.ProcessingCode, transactionAmount int64, transactionCurrency int32,
		settlementAmount int64, settlementCurrency int32, tokenId int64, mcc int64, bankData string) (string, error)

	Authorize(ctx context.Context, transactionId string) iso8583.ResponseCode

	GetTokenByHash(ctx context.Context, hash string) (int64, error)

	AdviceTransaction(ctx context.Context, reference string, extId string, trType iso8583.ProcessingCode, transactionAmount int64, transactionCurrency int32,
		settlementAmount int64, settlementCurrency int32, tokenId int64, mcc int64, bankData string, responseCode iso8583.ResponseCode) iso8583.ResponseCode
}

type GrpcClient struct {
	tokenConn      vaultv1.TokenServiceClient
	processingConn processingv1.TransactionServiceClient
}

func Init() error {
	conn, err := grpc.NewClient(config.Settings.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	GatekeeperEngine = &GrpcClient{
		tokenConn:      vaultv1.NewTokenServiceClient(conn),
		processingConn: processingv1.NewTransactionServiceClient(conn),
	}
	return nil
}
