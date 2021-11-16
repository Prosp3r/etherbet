package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func failOnError(err error, context string) {
	if err != nil {
		log.Fatalf("Failed %v with error : %v \n", context, err)
	}
}

func GenPrivatekey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		failOnError(err, "Generating Private key")
		errorMsg := "Failed generating private key"
		return nil, errors.New(errorMsg)
	}

	privateData := crypto.FromECDSA(privateKey)

	publicKey := GenPublicKey(privateKey)
	publicData := crypto.FromECDSAPub(publicKey)
	addressString := GenAddress(publicKey, privateKey)

	fmt.Printf("Public key: %v\n Private Key: %v \n Address: %v", hexutil.Encode(publicData), hexutil.Encode(privateData), addressString)

	// return privateKey, nil
	return nil, nil

}

func GenPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privateKey.PublicKey
}

func GenAddress(publicKey *ecdsa.PublicKey, privatekey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(privatekey.PublicKey).Hex()
}
