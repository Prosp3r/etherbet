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
	// denomination := "eth"
	denomination := "wei"
	accountEthBalance, err := getEthBalance(client, "0x737ec93DC736344a8A8EF51C337e052d29F61493")
	failOnError(err, "Getting balance")

	fmt.Printf("Received amount in %v : %v", denomination, accountEthBalance)



}


func getBalance(){

}

//getBalance takes in hex and denomination and returns amount on address in given hex in the denomination
func getEthBalance(client *ethclient.Client, hex string) (*big.Float, error) {
	denomination := "eth"
	address := common.HexToAddress(hex)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	failOnError(err, "Getting BalanceAt")

	if denomination == "eth" {
		//1 eht = 10 ^ 18 wei
		fmt.Println("Raw Balance in wei : ", balance)
		floatBalance := new(big.Float)
		floatBalance.SetString(balance.String())
		fmt.Printf("FloatBalance in %v : %v \n", denomination, floatBalance)
		balanceEther := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
		return balanceEther, nil
	}
	msg := "Could not get Balance"
	return nil, errors.New(msg)
}

//getBalance takes in hex and denomination and returns amount on address in given hex in the denomination
func getWeiBalance(client *ethclient.Client, hex, denomination string) *big.Int {
	address := common.HexToAddress(hex)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	failOnError(err, "Getting BalanceAt")
	//return as is
	fmt.Println("Raw Balance in wei : ", balance)
	return balance
}