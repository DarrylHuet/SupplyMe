package repo

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Balance() {
	client, err := ethclient.Dial("http://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	blocknumber := big.NewInt(5532883)
	balanceAt, err := client.BalanceAt(context.Background(), account, blocknumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balanceAt)
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
