package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func main() {
	logger := shim.NewLogger("user")
	userChaincode := chaincode.NewUserChaincode(logger)

	if err := shim.Start(userChaincode); err != nil {
		logger.Critical("Chaincode error: %s", err.Error())
	}
}
