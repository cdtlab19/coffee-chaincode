package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cdtlab19/coffee-chaincode/model"
	. "github.com/cdtlab19/coffee-chaincode/model"
)

var _ = Describe("User", func() {
	It("Should create a valid user", func() {
		user := NewUser("id", "someone", 3)
		Expect(user.DocType).To(Equal(UserDocType))
		Expect(user.ID).To(Equal("id"))
		Expect(user.Name).To(Equal("someone"))
		Expect(user.RemainingCoffee).To(Equal(3))

		err := user.Valid()
		Expect(err).NotTo(HaveOccurred())

	})

	It("Should have a valid ID", func() {
		user := NewUser("", "someone", 3)
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "", 3)
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "someone", 3)
		user.DocType = ""
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should drink a coffee", func() {
		user := NewUser("id", "someone", 3)
		err := user.DrinkCoffee()
		Expect(err).NotTo(HaveOccurred())
	})

	It("Should drink only if there's a coffee available", func() {
		user := NewUser("id", "someone", 0)

		err := user.DrinkCoffee()
		Expect(err).To(HaveOccurred())
	})

	It("Should be encodable", func() {
		jsonUser := NewUser("id", "someone", 3).JSON()

		var user model.User
		Expect(json.Unmarshal(jsonUser, &user)).NotTo(HaveOccurred())
		Expect(user.ID).To(Equal("id"))
		Expect(user.Name).To(Equal("someone"))
		Expect(user.RemainingCoffee).To(Equal(3))
	})
})
