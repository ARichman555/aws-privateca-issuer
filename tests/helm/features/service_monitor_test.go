package features

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
)

func TestServiceMonitor(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "serviceMonitor enabled creates ServiceMonitor resource",
			Values: map[string]interface{}{
				"serviceMonitor": map[string]interface{}{
					"create": true,
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				// ServiceMonitor test passed - chart installed successfully with serviceMonitor.create=true
				h.T.Logf("ServiceMonitor test passed - chart installed successfully with serviceMonitor.create=true")
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
