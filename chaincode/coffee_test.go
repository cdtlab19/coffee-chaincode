package chaincode_test

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cdtlab19/coffee-chaincode/chaincode"
	"github.com/cdtlab19/coffee-chaincode/model"
	"github.com/cdtlab19/coffee-chaincode/store"
)

var _ = Describe("Coffee", func() {
	var mock *shim.MockStub
	var logger *shim.ChaincodeLogger
	var st *store.CoffeeStore

	BeforeEach(func() {
		logger = shim.NewLogger("coffee-test")
		mock = shim.NewMockStub("coffee", NewCoffeeChaincode(logger))
		st = store.NewCoffeeStore(mock, logger)
	})

	It("Should Init", func() {
		result := mock.MockInit("0000", [][]byte{})
		Expect(int(result.Status)).To(Equal(shim.OK))
		Expect(result.Payload).To(BeEmpty())
	})

	It("Should CreateCoffee", func() {
		result := mock.MockInvoke("0000", [][]byte{
			[]byte("CreateCoffee"),
			[]byte("cappuccino"),
		})

		Expect(int(result.Status)).To(Equal(shim.OK))
		Expect(result.Payload).To(BeEmpty())

		coffee, err := st.GetCoffee("0000")
		Expect(err).NotTo(HaveOccurred())
		Expect(coffee.Flavour).To(Equal("cappuccino"))
	})

	Context("Method UseCoffee", func() {
		const method = "UseCoffee"

		It("Should return invalid if called with more than one argument", func() {
			// invoke method with 0 arguments
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

			// invoke method with 2 arguments
			result = mock.MockInvoke("0001", [][]byte{
				[]byte(method),
				[]byte("a"),
				[]byte("b"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))
		})

		It("Should execute successfuly", func() {
			// create asset for testing
			createTestCoffee(mock, st, model.NewCoffee("0000", "capuccino"))

			// invoke UseCoffee
			result := mock.MockInvoke("0000", [][]byte{
				[]byte("UseCoffee"),
				[]byte("0000"),
				[]byte("test-owner"),
			})

			// test if transaction was successful
			Expect(int(result.Status)).To(Equal(shim.OK))
			Expect(result.Payload).To(BeEmpty())

			// test if state changed
			coffee, err := store.NewCoffeeStore(mock, logger).GetCoffee("0000")
			Expect(err).NotTo(HaveOccurred())
			Expect(coffee.Owner).To(Equal("test-owner"))
		})
	})

	Context("DeleteCoffee", func() {
		const method = "DeleteCoffee"

		It("Should return invalid if called with more than one argument", func() {
			// invoke method with 0 arguments
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))

			// invoke method with 2 arguments
			result = mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("a"),
				[]byte("b"),
			})
			Expect(int(result.Status)).To(Equal(shim.ERROR))
		})

		It("Should execute successfuly", func() {
			createTestCoffee(mock, st, model.NewCoffee("0000", "capuccino"))

			// invoke UseCoffee
			result := mock.MockInvoke("0000", [][]byte{
				[]byte(method),
				[]byte("0000"),
			})

			// test if transaction was successful
			Expect(int(result.Status)).To(Equal(shim.OK))
			Expect(result.Payload).To(BeEmpty())

			// test if state changed
			_, err := store.NewCoffeeStore(mock, logger).GetCoffee("0000")
			Expect(err).To(HaveOccurred())
		})
	})

	It("Should DeleteCoffee", func() {
		// create asset for testing
		createTestCoffee(mock, st, model.NewCoffee("0000", "capuccino"))

		// invoke UseCoffee
		result := mock.MockInvoke("0000", [][]byte{
			[]byte("DeleteCoffee"),
			[]byte("0000"),
		})

		// test if transaction was successful
		Expect(int(result.Status)).To(Equal(shim.OK))
		Expect(result.Payload).To(BeEmpty())

		// test if state changed
		_, err := store.NewCoffeeStore(mock, logger).GetCoffee("0000")
		Expect(err).To(HaveOccurred())
	})
})

func createTestCoffee(mock *shim.MockStub, st *store.CoffeeStore, coffee *model.Coffee) {
	mock.MockTransactionStart("mocked")
	defer mock.MockTransactionEnd("mocked")

	if err := st.SetCoffee(coffee); err != nil {
		panic(err)
	}
}
