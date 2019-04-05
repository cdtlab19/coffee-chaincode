package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/cdtlab19/coffee-chaincode/store"
	"github.com/cdtlab19/coffee-chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CoffeeChaincode is a chaincode for controller coffee assets
type CoffeeChaincode struct {
	logger *shim.ChaincodeLogger
	router *utils.Router
}

// Checagem em tempo de compilação se CoffeeChaincode implementa shim.CoffeeChaincode
var _ shim.Chaincode = &CoffeeChaincode{}

// NewCoffeeChaincode cria uma nova instância do CoffeeChaincode para gerenciamento de
// cafés com os parâmetros default
func NewCoffeeChaincode(logger *shim.ChaincodeLogger) *CoffeeChaincode {
	chaincode := &CoffeeChaincode{logger: logger}
	chaincode.router = utils.NewRouter().
		Add("CreateCoffee", chaincode.CreateCoffee).
		Add("UseCoffee", chaincode.UseCoffee).
		Add("GetCoffee", chaincode.GetCoffee).
		Add("AllCoffee", chaincode.AllCoffee).
		Add("DeleteCoffee", chaincode.DeleteCoffee)

	return chaincode
}

// Init realiza as operações de inicialização do CoffeeChaincode
func (c *CoffeeChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke é chamado toda vez que o Chaicode é invocado
func (c *CoffeeChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	return c.router.Handle(stub, fn, args)
}

func (c *CoffeeChaincode) store(stub shim.ChaincodeStubInterface) *store.CoffeeStore {
	return store.NewCoffeeStore(stub, c.logger)
}

// CreateCoffee cria um novo café
func (c *CoffeeChaincode) CreateCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	coffee := &model.Coffee{}
	if err := json.Unmarshal([]byte(args[0]), &coffee); err != nil {
		return nil, err
	}

	if err := c.store(stub).SetCoffee(coffee); err != nil {
		return nil, err
	}

	return nil, nil
}

// UseCoffee uses a coffee capsule
func (c *CoffeeChaincode) UseCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 2 {
		return nil, errors.New("Wrong number of arguments")
	}

	st := c.store(stub)
	coffee, err := st.GetCoffee(args[0])
	if err != nil {
		return nil, err
	}

	if err := coffee.SetOwner(args[1]); err != nil {
		return nil, err
	}

	return nil, st.SetCoffee(coffee)
}

// GetCoffee retorna um café
func (c *CoffeeChaincode) GetCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	coffee, err := c.store(stub).GetCoffee(args[0])
	if err != nil {
		return nil, err
	}

	return struct {
		Coffee *model.Coffee `json:"coffee"`
	}{coffee}, nil
}

// AllCoffee retorna todos os cafés
func (c *CoffeeChaincode) AllCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("Função não recebe argumentos")
	}

	coffees, err := c.store(stub).AllCoffee()
	if err != nil {
		return nil, err
	}

	return struct {
		Coffees []*model.Coffee `json:"coffees"`
	}{coffees}, nil
}

// DeleteCoffee retorna todos os cafés
func (c *CoffeeChaincode) DeleteCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	return nil, c.store(stub).DeleteCoffee(args[0])
}
