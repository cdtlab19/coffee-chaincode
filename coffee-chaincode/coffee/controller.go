package coffee

import (
	"encoding/json"

	"github.com/cdtlab19/coffee-chaincode/base"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CoffeeChaincode implements a CRUD coffee CoffeeChaincode
type CoffeeChaincode struct {
	logger base.Logger
	router *base.Router
}

// Checagem em tempo de compilação se CoffeeChaincode implementa shim.CoffeeChaincode
var _ shim.Chaincode = &CoffeeChaincode{}

// NewCoffeeChaincode cria uma nova instância do CoffeeChaincode para gerenciamento de
// cafés com os parâmetros default
func NewCoffeeChaincode(logger base.Logger) *CoffeeChaincode {
	chaincode := &CoffeeChaincode{logger: logger}
	chaincode.router = base.NewRouter().
		Add("CreateCoffee", chaincode.CreateCoffee).
		Add("UpdateCoffee", chaincode.UpdateCoffee).
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

// Repository returns a coffee repository for this stub
func (c *CoffeeChaincode) Repository(stub shim.ChaincodeStubInterface) *CoffeeRepository {
	return NewCoffeeRepository(stub)
}

// CreateCoffee cria um novo café
func (c *CoffeeChaincode) CreateCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, base.NewValidationError("Precisa de '%' argumentos, recebido '%d'", 1, len(args))
	}

	coffee := &Coffee{}
	if err := json.Unmarshal([]byte(args[0]), &coffee); err != nil {
		return nil, err
	}

	if err := c.Repository(stub).SetCoffee(coffee); err != nil {
		return nil, err
	}

	return nil, nil
}

// UpdateCoffee atualiza um café
func (c *CoffeeChaincode) UpdateCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	return c.CreateCoffee(stub, args)
}

// GetCoffee retorna um café
func (c *CoffeeChaincode) GetCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, base.NewValidationError("Precisa de '%' argumentos, recebido '%d'", 1, len(args))
	}

	coffee, err := c.Repository(stub).GetCoffee(args[0])
	if err != nil {
		return nil, err
	}

	return struct {
		Coffee *Coffee `json:"coffee"`
	}{coffee}, nil
}

// AllCoffee retorna todos os cafés
func (c *CoffeeChaincode) AllCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 0 {
		return nil, base.NewValidationError("Função não recebe argumentos")
	}

	coffees, err := c.Repository(stub).AllCoffee()
	if err != nil {
		return nil, err
	}

	return struct {
		Coffees []*Coffee `json:"coffees"`
	}{coffees}, nil
}

// DeleteCoffee retorna todos os cafés
func (c *CoffeeChaincode) DeleteCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, base.NewValidationError("Precisa de '%' argumentos, recebido '%d'", 1, len(args))
	}

	return nil, c.Repository(stub).DeleteCoffee(args[0])
}
