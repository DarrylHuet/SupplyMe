package repo

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func CreateKeyStore() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

}

func importKeyStore() {
	file := ""
	ks := keystore.NewKeyStore("./tmp", keyStore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err ;= ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	
	password := ""
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(file), err != nil {
		log.Fatal(err)
	}
}
