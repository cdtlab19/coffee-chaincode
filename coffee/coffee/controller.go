package coffee

import (
	"encoding/json"

	"github.com/cdtlab19/coffee-chaincode/base"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode implementa as funções necessárias para execução via
// Hyperledger Fabric
type Chaincode struct {
	logger            *shim.ChaincodeLogger
	repositoryFactory RepositoryFactory
}

// Repository retorna um repository para um pedido
func (c *Chaincode) Repository(stub shim.ChaincodeStubInterface) Repository {
	return c.repositoryFactory(stub)
}

// Checagem em tempo de compilação se Chaincode implementa shim.Chaincode
var _ shim.Chaincode = &Chaincode{}

// NewChaincode cria uma nova instância do chaincode para gerenciamento de
// cafés com os parâmetros default
func NewChaincode(logger *shim.ChaincodeLogger, repositoryFactory RepositoryFactory) *Chaincode {
	return &Chaincode{
		logger:            logger,
		repositoryFactory: repositoryFactory,
	}
}

// Init realiza as operações de inicialização do Chaincode
func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke é chamado toda vez que o Chaicode é invocado
func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	var response interface{}
	var err error

	switch function {
	case "CreateCoffee":
		response, err = c.CreateCoffee(stub, args)
	case "UpdateCoffee":
		response, err = c.UpdateCoffee(stub, args)
	case "GetCoffee":
		response, err = c.GetCoffee(stub, args)
	case "AllCoffee":
		response, err = c.AllCoffee(stub, args)
	case "DeleteCoffee":
		response, err = c.DeleteCoffee(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	v, err := json.Marshal(response)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(v)
}

// CreateCoffee cria um novo café
func (c *Chaincode) CreateCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
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
func (c *Chaincode) UpdateCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	return c.CreateCoffee(stub, args)
}

// GetCoffee retorna um café
func (c *Chaincode) GetCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
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
func (c *Chaincode) AllCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
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
func (c *Chaincode) DeleteCoffee(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, base.NewValidationError("Precisa de '%' argumentos, recebido '%d'", 1, len(args))
	}

	return nil, c.Repository(stub).DeleteCoffee(args[0])
}
