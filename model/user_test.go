package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cdtlab19/coffee-chaincode/model"
)

var _ = Describe("User", func() {
	It("Should create a valid user", func() {
		user := NewUser("id", "someone")
		Expect(user.DocType).To(Equal(UserDocType))
		Expect(user.ID).To(Equal("id"))
		Expect(user.Name).To(Equal("someone"))

		err := user.Valid()
		Expect(err).NotTo(HaveOccurred())

	})

	It("Should have a valid ID", func() {
		user := NewUser("", "someone")
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "")
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should have a valid name", func() {
		user := NewUser("id", "someone")
		user.DocType = ""
		err := user.Valid()
		Expect(err).To(HaveOccurred())
	})

	It("Should be encodable", func() {
		user := NewUser("id", "someone")
		jsonUser := user.JSON()

		var raw map[string]string
		Expect(json.Unmarshal(jsonUser, &raw)).NotTo(HaveOccurred())
		Expect(raw).To(HaveKeyWithValue("id", user.ID))
		Expect(raw).To(HaveKeyWithValue("name", user.Name))
		Expect(raw).To(HaveKeyWithValue("docType", user.DocType))

	})

})
