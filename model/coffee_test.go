package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cdtlab19/coffee-chaincode/model"
)

var _ = Describe("Coffee", func() {
	It("should create a valid coffee", func() {
		coffee := NewCoffee("test", "capuccino")
		Expect(coffee.DocType).To(Equal(CoffeeDocType))
		Expect(coffee.Flavour).To(Equal("capuccino"))
		Expect(coffee.ID).To(Equal("test"))

		Expect(coffee.HasOwner()).To(Equal(false))

		err := coffee.Valid()
		Expect(err).NotTo(HaveOccurred())
	})

	It("should only set it's owner once", func() {
		var err error

		coffee := NewCoffee("test", "cappuccino")

		err = coffee.SetOwner("test")
		Expect(err).NotTo(HaveOccurred())
		Expect(coffee.Owner).To(Equal("test"))

		err = coffee.SetOwner("invalid")
		Expect(err).To(HaveOccurred())
		Expect(coffee.Owner).To(Equal("test"))
	})

	It("should have a valid ID", func() {
		coffee := NewCoffee("", "cappuccino")
		err := coffee.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("should have a valid DocType", func() {
		coffee := NewCoffee("id", "cappuccino")
		coffee.DocType = ""

		err := coffee.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("should be encodable", func() {
		coffee := NewCoffee("id", "chocolate")
		jsonObject := coffee.JSON()

		var raw map[string]string
		Expect(json.Unmarshal(jsonObject, &raw)).NotTo(HaveOccurred())
		Expect(raw).To(HaveKeyWithValue("id", coffee.ID))
		Expect(raw).To(HaveKeyWithValue("flavour", coffee.Flavour))
		Expect(raw).To(HaveKeyWithValue("docType", coffee.DocType))
		Expect(raw).To(HaveKeyWithValue("owner", coffee.Owner))
	})
})
