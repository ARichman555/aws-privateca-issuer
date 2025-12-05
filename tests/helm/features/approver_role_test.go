package features

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestApproverRole(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "approverRole enabled creates ClusterRole for certificate approval",
			Values: map[string]interface{}{
				"approverRole": map[string]interface{}{
					"enabled":            true,
					"namespace":          "cert-manager",
					"serviceAccountName": "cert-manager",
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				// Check for the approver ClusterRole
				approverRoleName := "cert-manager-controller-approve:awspca-cert-manager-io"
				clusterRole, err := h.Clientset.RbacV1().ClusterRoles().Get(context.TODO(), approverRoleName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, clusterRole)

				// Check for the approver ClusterRoleBinding
				clusterRoleBinding, err := h.Clientset.RbacV1().ClusterRoleBindings().Get(context.TODO(), approverRoleName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, clusterRoleBinding)
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
