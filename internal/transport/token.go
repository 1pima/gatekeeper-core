package transport

import (
	"context"
	vaultv1 "gatekeeper-core/pkg/proto/gatekeeper/contracts/vault/v1"
)

func (c *GrpcClient) GetTokenByHash(ctx context.Context, hash string) (int64, error) {
	resp, err := c.tokenConn.GetTokenByHash(ctx, &vaultv1.GetTokenByHashRequest{CardHash: hash})
	if err != nil {
		return 0, err
	}
	return resp.TokenId, nil
}
