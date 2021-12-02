package wallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var storage = "./wallet/ugodi"
var walletFiles []string
var addressDB map[string]string

func failOnError(err error, context string) {
	if err != nil {
		log.Fatalf("Failed %v with error : %v \n", context, err)
	}
}

func CreateAddress(password string) *accounts.Account {
	Keyx := keystore.NewKeyStore(storage, keystore.StandardScryptN, keystore.StandardScryptP)
	Newaddress, err := Keyx.NewAccount(password)
	failOnError(err, "Failed Creating New account")
	return &Newaddress
}

func ReadInAddresses() {
	err := filepath.Walk(storage, func(path string, info os.FileInfo, err error) error {
		walletFiles = append(walletFiles, path)
		return nil
	})
	failOnError(err, "Reading File names")

	for _, f := range walletFiles {
		fmt.Println(f)
		fileName := strings.Split(f, "/")
		keyName := strings.Split(fileName[2], "--")
		if len(keyName) > 2 {
			fmt.Printf("Keyname[0]: %v \v Keyname[1]: %v \n Keyname[2]: %v\n", keyName[0], keyName[1], keyName[2])
		}
		addy := keyName[2]
		addressDB[addy] = f
	}
	fmt.Printf("Total in address database : %v\n", len(addressDB))
}

func ReadMyAddress(password, address string) {
	// addys, err := ioutil.ReadFile(storage + "/")
	// failOnError(err, "Reading  Address Files")
	// for _, v := range addys {
	// 	fmt.Println(v)
	// }

}

func CreateKeys(password string) map[string]string {

	_ = CreateAddress(password)
	ReadInAddresses()

	ppkey := make(map[string]string)

	privateKey, err := GenPrivatekey()
	failOnError(err, "GeneratingPrivateKey")
	privateData := crypto.FromECDSA(privateKey)

	publicKey := GenPublicKey(privateKey)
	publicData := crypto.FromECDSAPub(publicKey)

	addressString := GenAddress(publicKey, privateKey)

	ppkey["private"] = hexutil.Encode(privateData)
	ppkey["public"] = hexutil.Encode(publicData)
	ppkey["address"] = addressString

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
