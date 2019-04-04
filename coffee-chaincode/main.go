package main

import (
	"github.com/cdtlab19/coffee-chaincode/coffee-chaincode/coffee"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	logger := shim.NewLogger("coffee")
	coffeeChaincode := coffee.NewCoffeeChaincode(logger)

	if err := shim.Start(coffeeChaincode); err != nil {
		logger.Critical("Chaincode Error: %s", err.Error())
	}
}
