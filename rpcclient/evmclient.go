package rpcclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmClient struct {
	ethClient *ethclient.Client
}

var _ RpcClient = (*EvmClient)(nil)

// NewEvmClient creates a new EvmClient instance.
func NewEvmClient(rpcURL string) (*EvmClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &EvmClient{
		ethClient: client,
	}, nil
}

// GetBlockHashByNumber returns the block hash by the given block number.
func (c *EvmClient) GetBlockHashByNumber(blockNumber int64) (string, error) {
	header, err := c.ethClient.HeaderByNumber(context.Background(), big.NewInt(blockNumber))
	return header.Hash().Hex(), err
}

// GetChainID returns the chain ID.
func (c *EvmClient) GetChainID() (int32, error) {
	chainID, err := c.ethClient.ChainID(context.Background())
	return int32(chainID.Int64()), err
}