package core

import (
	"context"
	"testing"

	"github.com/cert-manager/aws-privateca-issuer/tests/helm/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeployment(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "disableApprovedCheck adds command line flag",
			Values: map[string]interface{}{
				"disableApprovedCheck": true,
			},
			Validate: func(t *testing.T, h *testutil.TestHelper, releaseName string) {
				deploymentName := releaseName + "-aws-privateca-issuer"
				deployment, err := h.Clientset.AppsV1().Deployments(h.Namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
				require.NoError(t, err)

				container := deployment.Spec.Template.Spec.Containers[0]
				assert.Contains(t, container.Args, "-disable-approved-check")
			},
		},
	}

	testutil.RunTestCases(t, testCases)
}
