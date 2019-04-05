# Hyperledger Fabric Coffee Pod Manager

![image](https://travis-ci.com/vtfr/coffee-chain.svg?branch=master)

Coffee Pod Manager is a simple Blockchain application for learning purposes. It manages coffee pods by using Hyperledger Fabric's Distributed Ledger.

## Chaincodes

### Coffee Chaincode

The Chaincode `coffee' controlls usage of coffe pods

    github.com/cdtlab19/coffee-chaincode/entry/coffee

### Testing

    $ go get ./...
    $ go test ./...
