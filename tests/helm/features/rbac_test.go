package features

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRBAC(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "rbac enabled creates ClusterRole and ClusterRoleBinding",
			Values: map[string]interface{}{
				"rbac": map[string]interface{}{
					"create": true,
				},
				"serviceAccount": map[string]interface{}{
					"create": true,
				},
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				clusterRoleName := releaseName + "-aws-privateca-issuer"
				clusterRole, err := h.Clientset.RbacV1().ClusterRoles().Get(context.TODO(), clusterRoleName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, clusterRole)

				clusterRoleBinding, err := h.Clientset.RbacV1().ClusterRoleBindings().Get(context.TODO(), clusterRoleName, metav1.GetOptions{})
				require.NoError(t, err)
				assert.NotNil(t, clusterRoleBinding)
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
