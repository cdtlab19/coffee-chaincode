package chaincode

import (
	"fmt"

	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/cdtlab19/coffee-chaincode/store"
	"github.com/cdtlab19/coffee-chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// UserChaincode is a chaincode controller for user assets
type UserChaincode struct {
	logger *shim.ChaincodeLogger
	router *utils.Router
}

var _ shim.Chaincode = &UserChaincode{}

// NewUserChaincode cria uma nova instância do UserChaincode para gerenciamento de
// usuários com os parâmetros default
func NewUserChaincode(logger *shim.ChaincodeLogger) *UserChaincode {
	chaincode := &UserChaincode{logger: logger}
	chaincode.router = utils.NewRouter().
		Add("CreateUser", chaincode.CreateUser).
		Add("GetUser", chaincode.GetUser).
		Add("AllUser", chaincode.AllUser).
		Add("DeleteUser", chaincode.DeleteUser)

	return chaincode

}

// Init realiza as operações de inicialização do UserChaincode
func (u *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke é chamado toda vez que o Chaicode é invocado
func (u *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	return u.router.Handle(stub, fn, args)
}

func (u *UserChaincode) store(stub shim.ChaincodeStubInterface) *store.UserStore {
	return store.NewUserStore(stub, u.logger)
}

// CreateUser cria um novo usuário
func (u *UserChaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	user := model.NewUser(stub.GetTxID(), args[0])

	if err := u.store(stub).SetUser(user); err != nil {
		return nil, err
	}

	return nil, nil
}

// GetUser retorna um usuário
func (u *UserChaincode) GetUser(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	user, err := u.store(stub).GetUser(args[0])
	if err != nil {
		return nil, err
	}

	return struct {
		User *model.User `json:"user"`
	}{user}, nil
}

// AllUser retorna todos os usuários
func (u *UserChaincode) AllUser(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	users, err := u.store(stub).AllUser()
	if err != nil {
		return nil, err
	}

	return struct {
		Users []*model.User `json:"users"`
	}{users}, nil

}

// DeleteUser deleta um usuário
func (u *UserChaincode) DeleteUser(stub shim.ChaincodeStubInterface, args []string) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Precisa de '%d' argumentos, recebido '%d'", 1, len(args))
	}

	return nil, u.store(stub).DeleteUser(args[0])
}
