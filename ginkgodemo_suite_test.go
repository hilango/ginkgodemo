package ginkgodemo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgodemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgodemo Suite")
}
