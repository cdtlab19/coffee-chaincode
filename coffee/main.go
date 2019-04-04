package main

import (
	"github.com/cdtlab19/coffee-chaincode/coffee/coffee"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	logger := shim.NewLogger("coffee")
	coffeeChaincode := coffee.NewChaincode(logger, coffee.StubRepositoryFactory)

	if err := shim.Start(coffeeChaincode); err != nil {
		logger.Critical("Chaincode Error: %s", err.Error())
	}
}
