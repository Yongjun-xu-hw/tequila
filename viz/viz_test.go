package viz_test

import (
	. "github.com/newlee/tequila/viz"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Viz", func() {
	Context("Parse all include dependencies", func() {
		It("bc code", func() {
			codeDir := "../examples/bc-code/html"
			result := ParseCodeDir(codeDir)
			Expect(len(result.NodeList)).Should(Equal(12))

			crossRefs := result.FindCrossRef()

			Expect(len(crossRefs)).Should(Equal(0), "Cross references: %v", crossRefs)
		})
	})
})
