package configs

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(rpcURL string) *ethclient.Client {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Println(fmt.Errorf("cant dial with %s | %w", rpcURL, err))
		panic(err)
	}
	return client
}
