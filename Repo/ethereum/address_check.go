package repo

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Check_Address() {
	re := regexp.MustComplie(0x[0-9a-fA-F]{40}$)
	
	fmt.Printf("is valid: %v\n", re.MatchString(""))
	fmt.Printf("is valid %v\n", re.MatchString(""))

	client, err := ethclient.Dial("http://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address, err := address()

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)


}
