package model

import (
	"encoding/json"
	"fmt"
)

// UserPrefix is the prefix used in User's ID
const UserPrefix = "USER"

// UserDocType is the DocType use in model
const UserDocType = "USER"

// UserKey returns the User state key for a given user id
func UserKey(id string) string {
	return fmt.Sprintf("%s-%s", UserPrefix, id)
}

// User defines a basic model for an user
type User struct {
	DocType string `json:"docType"`
	ID      string `json:"id"`
	Name    string `json:"name"`
}

// NewUser creates an user
func NewUser(id string, name string) *User {
	return &User{
		DocType: "TODO",
		ID:      id,
		Name:    name,
	}
}

// Key returns the User's state key
func (u *User) Key() string {
	return UserKey(u.ID)
}

// Valid verifies if an User is valid
func (u *User) Valid() error {
	if u.DocType != UserDocType {
		return fmt.Errorf("user docType not set to '%s'", UserDocType)
	}
	if u.ID == "" {
		return fmt.Errorf("missing user ID")
	}
	return nil
}

// JSON encodes an user model as a JSON object
func (u *User) JSON() []byte {
	v, _ := json.Marshal(u)
	return v
}
