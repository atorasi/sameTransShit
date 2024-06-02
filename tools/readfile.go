package tools

import (
	"bufio"
	"dna/internal/configs"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func ReadStrings(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func SliceOfAccs(filepath string) ([]configs.Wallet, error) {
	keys, err := ReadStrings(filepath)
	if err != nil {
		return nil, err
	}

	accounts, err := wallets(keys)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func wallets(pks []string) ([]configs.Wallet, error) {
	var accounts []configs.Wallet
	for _, pk := range pks {
		account := newAccount(pk)

		accounts = append(accounts, account)
	}
	return accounts, nil
}

func newAccount(private string) configs.Wallet {
	if private[0:2] == "0x" {
		private = private[2:]
	}
	privateKey, _ := crypto.HexToECDSA(private)

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	return configs.Wallet{
		PrivateKey: privateKey,
		Address:    publicKey,
	}
}
