package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// UserDocType is the DocType use in model
const UserDocType = "user"

// User defines a basic model for an user
type User struct {
	DocType         string `json:"docType"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	RemainingCoffee string `json:"remainingCoffee"`
}

// NewUser creates an user with a exact amount of remaining coffees
func NewUser(id string, name string, remainingCoffee string) *User {
	return &User{
		DocType:         UserDocType,
		ID:              id,
		Name:            name,
		RemainingCoffee: remainingCoffee,
	}
}

// DrinkCoffee takes one unit of user's remaining coffees
func (u *User) DrinkCoffee() error {

	remainingCoffees, _ := strconv.Atoi(u.RemainingCoffee)
	if remainingCoffees < 1 {
		return errors.New("user has no remaining coffees")
	}

	u.RemainingCoffee = strconv.Itoa(remainingCoffees - 1)
	return nil
}

// Valid verifies if an User is valid
func (u *User) Valid() error {
	if u.DocType != UserDocType {
		return fmt.Errorf("user docType not set to '%s'", UserDocType)
	}
	if u.ID == "" {
		return fmt.Errorf("missing user ID")
	}

	if u.Name == "" {
		return fmt.Errorf("missing user name")
	}

	if _, err := strconv.Atoi(u.RemainingCoffee); err != nil {
		return fmt.Errorf("remaining coffee value is not valid; Expects an int string")
	}

	return nil
}

// JSON encodes an user model as a JSON object
func (u *User) JSON() []byte {
	v, _ := json.Marshal(u)
	return v
}
