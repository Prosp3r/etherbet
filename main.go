package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bitherhq/go-bither/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/8c5b190b405041f4afb69b99b46c4070"

func failOnError(err error, context string) {
	if err != nil {
		log.Fatalf("Failed %v with error %v", context, err)
	}
}

func main() {
	// ethclient.DialContext()
	client, err := ethclient.Dial(infuraURL)
	failOnError(err, "creating ether client")

	// defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	failOnError(err, "Getting BlockByNumber")
	fmt.Println(block.Number())
}