package repo

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infure.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	_ = client
}
