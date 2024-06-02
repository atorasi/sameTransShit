package transaction

import (
	"context"
	"fmt"
	"sameTrans/internal/configs"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendTransaction(client *ethclient.Client, wallet configs.Wallet, txData TransactionData) (common.Hash, error) {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("getting chain id %w", err)
	}
	transaction, err := prepareLegacy(client, wallet, txData)
	if err != nil {
		return common.Hash{}, fmt.Errorf("prepare transaction %w", err)
	}

	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainID), wallet.PrivateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("signing transaction %w", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("sending transaction %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	_, err = bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("wait for tx %w", err)
	}

	return signedTx.Hash(), nil
}

func prepareLegacy(client *ethclient.Client, wallet configs.Wallet, txData TransactionData) (*types.Transaction, error) {
	contractAddress := common.HexToAddress(txData.Contract)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("suggest gas %w", err)
	}

	message := ethereum.CallMsg{
		From:  wallet.Address,
		To:    &contractAddress,
		Data:  txData.Calldata,
		Value: txData.Value,
	}

	gasLimit, err := client.EstimateGas(context.Background(), message)
	if err != nil {
		return nil, fmt.Errorf("estimate gas %w", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), wallet.Address)
	if err != nil {
		return nil, fmt.Errorf("getting nonce %w", err)
	}

	return types.NewTransaction(nonce, contractAddress, txData.Value, gasLimit, gasPrice, txData.Calldata), nil
}
