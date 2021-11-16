package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/bitherhq/go-bither/common"
	"github.com/bitherhq/go-bither/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"
var ganacheURL = "http://127.0.0.1:8545"
var blockChain = infuraURL

func failOnError(err error, context string) {
	if err != nil {
		log.Fatalf("Failed %v with error %v", context, err)
	}
}

func main() {

	client, err := ethclient.Dial(blockChain)
	failOnError(err, "creating ether client")

	// defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	failOnError(err, "Getting BlockByNumber")
	fmt.Println("Block Number : ", block.Number())
	var denomination string

	denomination = "Eth"
	accountEthBalance, err := getEthBalance(client, "0x737ec93DC736344a8A8EF51C337e052d29F61493")
	failOnError(err, "Getting balance")
	fmt.Printf("Received amount in %v : %v", denomination, accountEthBalance)

	denomination = "wei"
	accountWeiBalance := getWeiBalance(client, "0x737ec93DC736344a8A8EF51C337e052d29F61493")
	fmt.Printf("Received amount in %v : %v", denomination, accountWeiBalance)

}

//getBalance - returns balance from blockchain
func getBalance(client *ethclient.Client, address common.Address) (*big.Int, error) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	// failOnError(err, "Getting BalanceAt")
	if err != nil {
		failOnError(err, "Could not get BalancAt")
		errorMsg := "Could not fetch balance from address"
		return nil, errors.New(errorMsg)
	}

	return balance, nil
}

//getEthBalance - takes in client, hex and returns amount on address in given hex in the denomination
func getEthBalance(client *ethclient.Client, hex string) (*big.Float, error) {
	address := common.HexToAddress(hex)
	balance, err := getBalance(client, address)
	failOnError(err, "Could not getBalance")

	fmt.Println("Raw Balance in wei : ", balance)
	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	fmt.Printf("FloatBalance in ETH : %v \n", floatBalance)
	balanceEther := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	return balanceEther, nil
}

//getBalance takes in hex and denomination and returns amount on address in given hex in the denomination
func getWeiBalance(client *ethclient.Client, hex string) *big.Int {
	address := common.HexToAddress(hex)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	failOnError(err, "Getting BalanceAt")
	//return as is
	fmt.Println("Raw Balance in wei : ", balance)
	return balance
}
