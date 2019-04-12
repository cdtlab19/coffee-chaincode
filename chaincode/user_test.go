package chaincode_test

import (
	"encoding/json"
	"fmt"

	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/cdtlab19/coffee-chaincode/store"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cdtlab19/coffee-chaincode/chaincode"
)

var _ = Describe("User", func() {
	var mock *shim.MockStub
	var logger *shim.ChaincodeLogger
	var st *store.UserStore

	BeforeEach(func() {
		logger = shim.NewLogger("user-test")
		mock = shim.NewMockStub("user", NewUserChaincode(logger))
		st = store.NewUserStore(mock, logger)
	})

	It("Should Init", func() {
		result := mock.MockInit("0000", [][]byte{})
		Expect(int(result.Status)).To(Equal(shim.OK))
		Expect(result.Payload).To(BeEmpty())
	})

	Context("CreateUser method", func() {
		const method = "CreateUser"

		It("Shoud create an user", func() {
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("name"),
			})

			Expect(int(result.Status)).To(Equal(shim.OK))

			// verify payload
			var response struct {
				User *model.User `json:"user"`
			}

			Expect(json.Unmarshal(result.Payload, &response)).ToNot(HaveOccurred())
			Expect(response.User.Name).To(Equal("name"))
			Expect(response.User.ID).To(Equal("0000"))

			user, err := st.GetUser("0000")
			Expect(err).NotTo(HaveOccurred())
			Expect(user.Name).To(Equal("name"))
		})
	})

	Context("GetUser", func() {
		const method = "GetUser"

		It("Should return error if no user was found", func() {
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("0000"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))
			Expect(result.Payload).To(BeEmpty())
		})

		It("Should return an user if it exists", func() {
			createTestUser(mock, st, model.NewUser("0000", "someone"))

			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("0000"),
			})
			Expect(int(result.Status)).To(Equal(shim.OK))
			Expect(result.Payload).NotTo(BeEmpty())

			// Payload: { "user": {...} }
			var response struct {
				User *model.User `json:"user"`
			}

			Expect(json.Unmarshal(result.Payload, &response)).NotTo(HaveOccurred())
			Expect(response.User.ID).To(Equal("0000"))
			Expect(response.User.Name).To(Equal("someone"))

		})
	})

	Context("AllUser", func() {
		const method = "AllUser"
		It("Should return all users", func() {
			user1 := model.NewUser("0000", "Someone")
			user2 := model.NewUser("0001", "Anyone")

			createTestUser(mock, st, user1)
			createTestUser(mock, st, user2)

			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
			})

			Expect(int(result.Status)).To(Equal(shim.OK))

			var res struct {
				Users []*model.User `json:"users"`
			}

			Expect(json.Unmarshal(result.Payload, &res)).ToNot(HaveOccurred())

			fmt.Printf("%+v", res)

			Expect(res.Users).To(HaveLen(2))
			Expect(res.Users).To(ContainElement(user1))
			Expect(res.Users).To(ContainElement(user2))
		})

	})

	Context("DeleteUser", func() {
		const method = "DeleteUser"

		It("Shoud return invalid if called with something other than one argument", func() {
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

			result = mock.MockInvoke("0001", [][]byte{
				[]byte(method),
				[]byte("a"),
				[]byte("b"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

		})

		It("Should WHAT", func() {
			createTestUser(mock, st, model.NewUser("0000", "someone"))

			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("0000"),
			})

			Expect(int(result.Status)).To(Equal(shim.OK))
			Expect(result.Payload).To(BeEmpty())

		})

	})

})

func createTestUser(mock *shim.MockStub, st *store.UserStore, user *model.User) {
	mock.MockTransactionStart("int")
	defer mock.MockTransactionEnd("int")

	if err := st.SetUser(user); err != nil {
		panic(err)
	}
}
