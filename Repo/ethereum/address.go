package repo

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func address() {
	address := common.HexToAddress("")

	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
