package chaincode_test

import (
	"encoding/json"

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
			Expect(result.Payload).To(BeEmpty())

			user, err := st.GetUser("0000")
			Expect(err).NotTo(HaveOccurred())
			Expect(user.ID).To(Equal("0000"))
			Expect(user.Name).To(Equal("name"))
		})

		It("Should return invalid if called with something other than one argument", func() {
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
		/*
			It("Should return invalid if called with something other than one argument", func() {
				result := mock.MockInvoke("0000", [][]byte{
					[]byte(method),
					[]byte(""),
				})

				Expect(int(result.Status)).To(Equal(shim.ERROR))
			})
		*/
	})

	Context("GetUser", func() {
		const method = "GetUser"

		It("Expects nothing different than 1 argument", func() {
			//invoking with 0 arguments
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

			//invoking with more than 1 argument
			result = mock.MockInvoke("0001", [][]byte{
				[]byte(method),
				[]byte("a"),
				[]byte("b"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

		})

		It("Should return error if no user was found", func() {
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("anythingElse"),
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

		It("Should return an error if called with any argument", func() {
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("something"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))
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
	mock.MockTransactionStart("mocked")
	defer mock.MockTransactionEnd("mocked")

	if err := st.SetUser(user); err != nil {
		panic(err)
	}
}
