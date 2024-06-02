package main

import (
	"dna/internal/app"
	"dna/internal/constants"
	"dna/tools"
	"fmt"
)

func main() {
	fmt.Println(constants.LOGO)

	wallets, err := tools.SliceOfAccs("./configs/private.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(wallets))

	for index, wallet := range wallets {
		fmt.Printf("%v | Start\n", index)
		if constants.SETTINGS.NeedZerion {
			client := constants.Clients["ETHEREUM"]
			if err := app.MintZerion(client, wallet); err != nil {
				fmt.Printf("%v | Error: %s", index, err)
			}
		}

		client := constants.Clients["POLYGON"]
		if err = app.RunPolygon(client, wallet); err != nil {
			fmt.Printf("%v | Error: %s\n", index, err)
		}

	}
	fmt.Println("Thanks for using script")

}
