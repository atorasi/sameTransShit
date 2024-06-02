package app

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sameTrans/internal/configs"
	"sameTrans/internal/constants"
	"sameTrans/internal/transaction"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Run(client *ethclient.Client, wallet configs.Wallet) error {
	calldata, _ := hex.DecodeString(constants.SETTINGS.INPUTDATA)

	txData := transaction.TransactionData{
		Value:    big.NewInt(constants.SETTINGS.VALUE),
		Contract: constants.SETTINGS.CONTRACT,
		Calldata: calldata,
	}

	txHash, err := transaction.SendTransaction(client, wallet, txData)
	if err != nil {
		return err
	}
	fmt.Printf("%s | %s%s\n", wallet.Address, constants.SETTINGS.SCAN, txHash.Hex())

	return nil
}
