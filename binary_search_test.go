package binarysearch

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIndexOf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binary Search Test Suite")
}

var _ = Describe("Binary Search", func() {
	It("should find an element", func() {
		Expect(IndexOf([]int{1,2,3,4}, 3)).To(Equal(2))
	})
})
