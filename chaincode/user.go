package chaincode

import (
	"github.com/vtfr/rocha/argsmw"

	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/cdtlab19/coffee-chaincode/store"
	"github.com/cdtlab19/coffee-chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/vtfr/rocha"
)

// UserChaincode is a chaincode controller for user assets
type UserChaincode struct {
	logger *shim.ChaincodeLogger
	router *rocha.Router
}

var _ shim.Chaincode = &UserChaincode{}

// NewUserChaincode cria uma nova instância do UserChaincode para gerenciamento de
// usuários com os parâmetros default
func NewUserChaincode(logger *shim.ChaincodeLogger) *UserChaincode {
	chaincode := &UserChaincode{logger: logger}
	chaincode.router = rocha.NewRouter().
		// CreateUser creates a new user with a `name`
		Handle("CreateUser",
			utils.RespondJSON(chaincode.CreateUser),
			argsmw.Arguments(argsmw.String("name"))).
		// GetUser returns an user by it's id
		Handle("GetUser", utils.RespondJSON(chaincode.GetUser),
			argsmw.Arguments(argsmw.String("id"))).
		// AllUser returns all users
		Handle("AllUser", utils.RespondJSON(chaincode.AllUser)).
		// DeleteUser deles an user by it's `id`
		Handle("DeleteUser", utils.RespondJSON(chaincode.DeleteUser),
			argsmw.Arguments(argsmw.String("id")))

	return chaincode

}

// Init realiza as operações de inicialização do UserChaincode
func (u *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke é chamado toda vez que o Chaicode é invocado
func (u *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	return u.router.Invoke(stub, fn, args)
}

// store
func (u *UserChaincode) store(stub shim.ChaincodeStubInterface) *store.UserStore {
	return store.NewUserStore(stub, u.logger)
}

// CreateUser cria um novo usuário
func (u *UserChaincode) CreateUser(c rocha.Context) (interface{}, error) {
	stub := c.Stub()

	user := model.NewUser(stub.GetTxID(), c.String("name"))

	if err := u.store(stub).SetUser(user); err != nil {
		return nil, err
	}

	return struct {
		User *model.User `json:"user"`
	}{user}, nil
}

// GetUser retorna um usuário
func (u *UserChaincode) GetUser(c rocha.Context) (interface{}, error) {
	user, err := u.store(c.Stub()).GetUser(c.String("id"))
	if err != nil {
		return nil, err
	}

	return struct {
		User *model.User `json:"user"`
	}{user}, nil
}

// AllUser retorna todos os usuários
func (u *UserChaincode) AllUser(c rocha.Context) (interface{}, error) {
	users, err := u.store(c.Stub()).AllUser()
	if err != nil {
		return nil, err
	}

	return struct {
		Users []*model.User `json:"users"`
	}{users}, nil
}

// DeleteUser deleta um usuário
func (u *UserChaincode) DeleteUser(c rocha.Context) (interface{}, error) {
	return nil, u.store(c.Stub()).DeleteUser(c.String("id"))
}
