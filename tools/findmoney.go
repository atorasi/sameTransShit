package tools

import (
	"context"
	"dna/internal/configs"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func MaxNative(clients []*ethclient.Client, wallet configs.Wallet) (*ethclient.Client, error) {
	var balance *big.Int
	var client *ethclient.Client

	ctx := context.Background()

	for _, cli := range clients {
		clientBalance, err := cli.BalanceAt(ctx, wallet.Address, nil)
		if err != nil {
			chainID, _ := cli.ChainID(ctx)
			fmt.Println(fmt.Errorf("cannot get balance at %v: %w", chainID, err))
		}
		if clientBalance.Cmp(balance) != -1 {
			balance = clientBalance
			client = cli
		}

	}
	return client, nil
}
