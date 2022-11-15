package self_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/Abdoueck632/mp-quic/integrationtests/tools/testlog"

	"testing"
)

func TestSelf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Self integration tests")
}
