package sys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKernelInfo_NoParamsReqd_HappyScenario(t *testing.T) {
	kernelInfo := GetKernelInfo()

	assert.NotEmpty(t, kernelInfo, "Expected non-empty kernel info string")
}

func TestGetHostname_NoParamsReqd_HappyScenario(t *testing.T) {
	hostname := GetHostname()

	assert.NotEmpty(t, hostname, "Expected non-empty hostname string")
}
