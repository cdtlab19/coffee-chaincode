package store

import (
	"encoding/json"

	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// UserStore abstracts user CRUD methods
type UserStore struct {
	stub   shim.ChaincodeStubInterface
	logger *shim.ChaincodeLogger
}

// NewUserStore creates a new user Store
func NewUserStore(stub shim.ChaincodeStubInterface, logger *shim.ChaincodeLogger) *UserStore {
	return &UserStore{stub, logger}
}

// AllUsers returns all existing users
func (u *UserStore) AllUsers() ([]*model.User, error) {
	u.logger.Debug("Entered AllUsers")
	iterator, err := u.stub.GetStateByPartialCompositeKey(model.UserDocType, []string{model.UserDocType})
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	u.logger.Debug("AllUsers: starting iterator")
	users := make([]*model.User, 1)
	for iterator.HasNext() {
		k, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		user := &model.User{}
		if err := json.Unmarshal(k.GetValue(), &user); err != nil {
			return nil, err
		}

		u.logger.Debugf("AllUsers: element with ID '%s' found", user.ID)
		users = append(users, user)
	}

	u.logger.Debug("Exiting AllUsers")
	return users, nil

}
