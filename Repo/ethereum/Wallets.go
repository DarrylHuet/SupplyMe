package repo

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"

)

func Error_Check(err) error {
	if err != nil {
		log.Fatal(err)
	} else 
	return nil
}

func CreateWallet() {
	privateKey, err := crypto.GenerateKey()
	Error_Check(err)
	
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println(hexutil.Encode(privateKeyBytes)[2:1])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

}

