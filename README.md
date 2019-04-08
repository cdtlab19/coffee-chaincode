# Hyperledger Fabric Coffee Pod Manager

[![Build Status](https://travis-ci.com/cdtlab19/coffee-chaincode.svg?branch=master)](https://travis-ci.com/cdtlab19/coffee-chaincode)
[![codecov](https://codecov.io/gh/cdtlab19/coffee-chaincode/branch/master/graph/badge.svg)](https://codecov.io/gh/cdtlab19/coffee-chaincode)
[![GoDoc](https://godoc.org/github.com/cdtlab19/coffee-chaincode?status.svg)](https://godoc.org/github.com/cdtlab19/coffee-chaincode)

Coffee Pod Manager is a simple Blockchain application for learning purposes. It manages coffee pods by using Hyperledger Fabric's Distributed Ledger.

## Chaincodes

### Coffee Chaincode

The Chaincode `coffee` controlls usage of coffe pods

    github.com/cdtlab19/coffee-chaincode/entry/coffee

### Testing

    $ go get -u -t ./...
    $ go test ./...
