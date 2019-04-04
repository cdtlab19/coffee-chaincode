package coffee

import (
	"encoding/json"
	"fmt"
	"strings"
)

const CoffeeDocType = "COFFEE"

// Coffee defines a basic model for coffee
type Coffee struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	Owner   string `json:"owner"`
}

// NewCoffee creates a new Coffee
func NewCoffee(id string, tp string) *Coffee {
	return &Coffee{
		DocType: CoffeeDocType,
		ID:      id,
		Type:    tp,
		Owner:   "",
	}
}

// HasOwner verifies if a Coffe has a owner
func (c *Coffee) HasOwner() bool {
	return c.Owner == ""
}

// IsValid verifies if a Coffee is valid or not
func (c *Coffee) IsValid() error {
	if c.DocType != CoffeeDocType {
		return fmt.Errorf("Coffee docType not set to '%s'", CoffeeDocType)
	}
	if !strings.HasPrefix(c.ID, CoffeeDocType) {
		return fmt.Errorf("Missing '%s' prefix from ID", CoffeeDocType)
	}
	return nil
}

// JSON encodifica um café como JSON
func (c *Coffee) JSON() []byte {
	v, _ := json.Marshal(c)
	return v
}
