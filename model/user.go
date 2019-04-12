package model

import (
	"encoding/json"
	"fmt"
)

// UserDocType is the DocType use in model
const UserDocType = "user"

// User defines a basic model for an user
type User struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`
	Name    string `json:"name"`
}

// NewUser creates an user
func NewUser(id string, name string) *User {
	return &User{
		DocType: UserDocType,
		ID:      id,
		Name:    name,
	}
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

	return nil
}

// JSON encodes an user model as a JSON object
func (u *User) JSON() []byte {
	v, _ := json.Marshal(u)
	return v
}
