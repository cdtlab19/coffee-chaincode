package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cdtlab19/coffee-chaincode/model"
)

var _ = Describe("User", func() {
	It("Should create a valid user", func() {
		user := NewUser("id", "someone", "3")
		Expect(user.DocType).To(Equal(UserDocType))
		Expect(user.ID).To(Equal("id"))
		Expect(user.Name).To(Equal("someone"))
		Expect(user.RemainingCoffee).To(Equal("3"))

		err := user.Valid()
		Expect(err).NotTo(HaveOccurred())

	})

	It("Should have a valid ID", func() {
		user := NewUser("", "someone", "3")
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "", "3")
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "someone", "3")
		user.DocType = ""
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid value of remaining coffees", func() {
		user := NewUser("id", "someone", "NaN")
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should drink a coffee", func() {
		user := NewUser("id", "someone", "3")
		err := user.DrinkCoffee()
		Expect(err).NotTo(HaveOccurred())
	})

	It("Should drink only if there's a coffee available", func() {
		user := NewUser("id", "someone", "3")
		user.RemainingCoffee = "0"

		err := user.DrinkCoffee()
		Expect(err).To(HaveOccurred())
	})

	It("Should be encodable", func() {
		user := NewUser("id", "someone", "3")
		jsonUser := user.JSON()

		var raw map[string]string
		Expect(json.Unmarshal(jsonUser, &raw)).NotTo(HaveOccurred())
		Expect(raw).To(HaveKeyWithValue("id", user.ID))
		Expect(raw).To(HaveKeyWithValue("name", user.Name))
		Expect(raw).To(HaveKeyWithValue("docType", user.DocType))
		Expect(raw).To(HaveKeyWithValue("remainingCoffee", user.RemainingCoffee))
	})
})
