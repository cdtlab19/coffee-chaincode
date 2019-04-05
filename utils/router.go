package utils

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// RouterFunc is a chaincode router function
type RouterFunc func(shim.ChaincodeStubInterface, []string) (interface{}, error)

// Router is a simple chaincode function router
type Router struct {
	routes map[string]RouterFunc
}

// NewRouter creates a new chaincode router
func NewRouter() *Router {
	return &Router{make(map[string]RouterFunc)}
}

// Add adds a new chaincode router function
func (r *Router) Add(name string, fn RouterFunc) *Router {
	r.routes[name] = fn
	return r
}

// Handle executes a router function or returns error if it was not found
func (r *Router) Handle(stub shim.ChaincodeStubInterface, name string, args []string) pb.Response {
	fn, ok := r.routes[name]
	if !ok {
		return shim.Error(fmt.Sprintf("Function '%s' not found in chaincode", name))
	}

	// executes the chaincode
	response, err := fn(stub, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error executing chaincode: %s", err.Error()))
	}

	if response == nil {
		return shim.Success(nil)
	}

	v, err := json.Marshal(response)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error encoding response: %s", err.Error()))
	}
	return shim.Success(v)
}
