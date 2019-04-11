package utils_test

import (
	"errors"
	"math"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"

	. "github.com/cdtlab19/coffee-chaincode/utils"
)

var _ = Describe("Respond", func() {

	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, "", []string{})
	})

	It("Should responde a valid JSON if no error is sent", func() {
		CONTENT := struct {
			Key1 string `json:"key1"`
			Key2 string `json:"key2"`
		}{"value1", "value2"}

		CONTENT_JSON := []byte(`{"key1":"value1","key2":"value2"}`)

		handler := RespondJSON(func(c rocha.Context) (interface{}, error) {
			return CONTENT, nil
		})

		resp := handler(context)
		Expect(int(resp.Status)).To(Equal(shim.OK))
		Expect(resp.Payload).To(Equal(CONTENT_JSON))
	})

	It("Should responde a empty response if empty data", func() {
		handler := RespondJSON(func(c rocha.Context) (interface{}, error) {
			return nil, nil
		})

		resp := handler(context)
		Expect(int(resp.Status)).To(Equal(shim.OK))
		Expect(resp.Payload).To(BeEmpty())
	})

	It("Should return an error if received an error", func() {
		ERROR_MESSAGE := "error message"

		handler := RespondJSON(func(c rocha.Context) (interface{}, error) {
			return nil, errors.New(ERROR_MESSAGE)
		})

		resp := handler(context)
		Expect(int(resp.Status)).To(Equal(shim.ERROR))
		Expect(resp.Message).To(Equal(ERROR_MESSAGE))
	})

	It("Should return an error if failed to encode JSON", func() {
		handler := RespondJSON(func(c rocha.Context) (interface{}, error) {
			// json can't encode math.Inf(1), so it will always return an error
			return math.Inf(1), nil
		})

		resp := handler(context)
		Expect(int(resp.Status)).To(Equal(shim.ERROR))
		Expect(resp.Message).To(ContainSubstring("Failed encoding response"))
	})
})
