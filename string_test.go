package types_test

import (
	"github.com/figoxu/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	It("IsEmail", func() {
		Ω(types.NewString("xujianhui@reyun.com").IsEmail()).Should(BeTrue())
		Ω(types.NewString("xujianhuireyun.com").IsEmail()).Should(BeFalse())
	})
})
