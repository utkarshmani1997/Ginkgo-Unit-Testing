package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSum Suite")
}
