package wallet

import (
	"crypto/ecdsa"
	"errors"
	"log"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func failOnError(err error, context string) {
	if err != nil {
		log.Fatalf("Failed %v with error : %v \n", context, err)
	}
}

func CreateKeys() map[string]string {

	ppkey := make(map[string]string)

	privateKey, err := GenPrivatekey()
	failOnError(err, "GeneratingPrivateKey")

	privateData := crypto.FromECDSA(privateKey)

	publicKey := GenPublicKey(privateKey)

	publicData := crypto.FromECDSAPub(publicKey)
	addressString := GenAddress(publicKey, privateKey)

	//fmt.Printf("Public key: %v\n Private Key: %v \n Address: %v", hexutil.Encode(publicData), hexutil.Encode(privateData), addressString)

	ppkey["private"] = hexutil.Encode(privateData)
	ppkey["public"] = hexutil.Encode(publicData)

	fmt.Printf("Private: %v \n Public: %v \n FirstKey: %v \n", ppkey["private"], ppkey["public"], addressString)
	addressString2 := GenAddress(publicKey, privateKey)
	fmt.Println("Second key: ", addressString2)
	return ppkey
}

func GenPrivatekey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		failOnError(err, "Generating Private key")
		errorMsg := "Failed generating private key"
		return nil, errors.New(errorMsg)
	}

	return privateKey, nil

}

func GenPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}

func GenAddress(publicKey *ecdsa.PublicKey, privatekey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(privatekey.PublicKey).Hex()
}
