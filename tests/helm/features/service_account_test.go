package features

import (
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
)

func TestServiceAccount(t *testing.T) {
	testCases := []testutil.PrivateCaHelmTestCase{
		{
			Name: "serviceAccount with custom name",
			Values: testutil.PrivateCaServiceAccountValues("custom-service-account", map[string]string{
				"eks.amazonaws.com/role-arn": "arn:aws:iam::123456789012:role/test-role",
			}),
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				serviceAccount := h.GetPrivateCaServiceAccount("custom-service-account")
				assert.Equal(t, "arn:aws:iam::123456789012:role/test-role", serviceAccount.Annotations["eks.amazonaws.com/role-arn"])
			},
		},
	}

	testutil.RunPrivateCaHelmTests(t, testCases)
}
