package main

import (
	"fmt"
	"math/rand"
	"sameTrans/internal/app"
	"sameTrans/internal/constants"
	"sameTrans/tools"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println(constants.LOGO)
	fmt.Println("-----------> t.me/tripleshizu | t.me/cryptosvinarnik <-----------")

	time.Sleep(time.Second)

	runScript()

	fmt.Println("\n\nThank you for using the software. </3")

}

func runScript() {
	client, err := ethclient.Dial(constants.SETTINGS.RPC)
	if err != nil {
		panic(err)
	}

	wallets, err := tools.SliceOfAccs("configs\\privatekeys.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wallets), "wallets loaded <-----------")

	for index, wallet := range wallets {
		fmt.Println(index, "Starting script on account:", wallet.Address)
		if err := app.Start(client, wallet); err != nil {
			fmt.Printf("%v | %s | Error: %s\n", index, wallet.Address, err)
		}
		if constants.SETTINGS.NEEDDELAYS {
			timeout := rand.Intn(constants.SETTINGS.SLEEPBETWEENACCOUNTWORK[1]+1) + constants.SETTINGS.SLEEPBETWEENACCOUNTWORK[0]

			fmt.Printf("%v | Sleeping %v seconds before next account\n", index, timeout)
			time.Sleep(time.Second * time.Duration(timeout))
		}
	}
}
