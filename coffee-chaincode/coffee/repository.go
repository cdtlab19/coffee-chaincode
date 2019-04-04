package coffee

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CoffeeRepository abstracts coffee repository methods
type CoffeeRepository struct {
	stub shim.ChaincodeStubInterface
}

// NewCoffeeRepository creates a new CoffeeRepository
func NewCoffeeRepository(stub shim.ChaincodeStubInterface) *CoffeeRepository {
	return &CoffeeRepository{stub}
}

// AllCoffee returns all existing coffee capsules
func (c *CoffeeRepository) AllCoffee() ([]*Coffee, error) {
	iterator, err := c.stub.GetStateByPartialCompositeKey(CoffeeDocType, []string{CoffeeDocType})
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	coffees := make([]*Coffee, 1)
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

// GetCoffee returns a coffee by it's id
func (c *CoffeeRepository) GetCoffee(coffeeID string) (coffee *Coffee, err error) {
	data, err := c.stub.GetState(coffeeID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &coffee)
	return
}

// SetCoffee sets a coffee asset by it's id
func (c *CoffeeRepository) SetCoffee(coffee *Coffee) error {
	return c.stub.PutState(coffee.ID, coffee.JSON())
}

// DeleteCoffee deletes a coffee asset by it's id
func (c *CoffeeRepository) DeleteCoffee(coffeeID string) error {
	return c.stub.DelState(coffeeID)
}
