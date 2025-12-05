package features

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestServiceAccount(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "serviceAccount with custom name",
			Values: map[string]interface{}{
				"serviceAccount": map[string]interface{}{
					"create": true,
					"name":   "custom-service-account",
					"annotations": map[string]interface{}{
						"eks.amazonaws.com/role-arn": "arn:aws:iam::123456789012:role/test-role",
					},
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				serviceAccount, err := h.Clientset.CoreV1().ServiceAccounts(h.Namespace).Get(context.TODO(), "custom-service-account", metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, serviceAccount)
				assert.Equal(t, "arn:aws:iam::123456789012:role/test-role", serviceAccount.Annotations["eks.amazonaws.com/role-arn"])
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
