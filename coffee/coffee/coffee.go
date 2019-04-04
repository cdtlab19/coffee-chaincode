package coffee

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CoffeeDocType contém o tipo de dado armazenado no ledger
const CoffeeDocType = "COFFEE"

// Coffe define o modelo básico de um café
type Coffee struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Owner string `json:"owner"`
}

// JSON encodifica um café como JSON
func (c *Coffee) JSON() []byte {
	v, _ := json.Marshal(c)
	return v
}

// Repository define um serviço de gerenciamento de cafés
//go:generate mockgen -destination=./mock_repository.go -package=coffee_test github.com/cdtlab19/coffee-chaincode/coffee/coffee MockRepository
type Repository interface {
	AllCoffee() ([]*Coffee, error)
	GetCoffee(coffeeID string) (*Coffee, error)
	SetCoffee(*Coffee) error
	DeleteCoffee(coffeeID string) error
}

// RepositoryFactory abstrai a criação de um repository, usado para que, durante o teste, seja possível utilizar
// um repository mockado
type RepositoryFactory func(shim.ChaincodeStubInterface) Repository

// StubRepositoryFactory cria um repositorio a partir das informações fornecidas pelo
// ChaincodeStubInterface
func StubRepositoryFactory(stub shim.ChaincodeStubInterface) Repository {
	return NewRepository(stub)
}

// MockRepositoryFactory retorna um Repository mockado toda vez que for executado, usado para
// testes da lógica de negócios
func MockRepositoryFactory(mock Repository) RepositoryFactory {
	return func(_ shim.ChaincodeStubInterface) Repository {
		return mock
	}
}

// repository é uma implementação do Repository
type repository struct {
	stub shim.ChaincodeStubInterface
}

// NewRepository cria um novo repositório de cafés
func NewRepository(stub shim.ChaincodeStubInterface) Repository {
	return &repository{stub}
}

// AllCoffee retorna todos os cafés
func (c *repository) AllCoffee() (coffees []*Coffee, err error) {
	iterator, err := c.stub.GetStateByPartialCompositeKey(CoffeeDocType, []string{CoffeeDocType})
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	coffees = make([]*Coffee, 1)
	for iterator.HasNext() {
		k, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		coffee := &Coffee{}
		if err := json.Unmarshal(k.GetValue(), &coffee); err != nil {
			return nil, err
		}

		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

// GetCoffee retorna um café
func (c *repository) GetCoffee(coffeeID string) (coffee *Coffee, err error) {
	data, err := c.stub.GetState(coffeeID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &coffee)
	return
}

// SetCoffee cria ou atualiza um café
func (c *repository) SetCoffee(coffee *Coffee) error {
	return c.stub.PutState(coffee.ID, coffee.JSON())
}

// DeleteCoffee deleta um café
func (c *repository) DeleteCoffee(coffeeID string) error {
	return c.stub.DelState(coffeeID)
}
