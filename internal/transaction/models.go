package transaction

import "math/big"

type TransactionData struct {
	Contract string
	Calldata []byte
	Value    *big.Int
}
